package processors

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/talbx/mqtt-wire/model"
	"github.com/talbx/mqtt-wire/utils"
	"time"
)

type SensorProcessor struct{}

var sensors = []string{prefix + "/sonoff-motion"}

func (processor SensorProcessor) Process(msg MQTT.Message, val string) []byte {
	var data model.SensorData
	err := json.Unmarshal(msg.Payload(), &data)
	utils.PrintIfErr(err)
	var record model.SensorRecord

	if val == "" {
		record = processor.createNewRecord(msg)
		utils.LogStr("Creating new record for topic", msg.Topic())
	} else {
		record = processor.updateExistingRecord(msg, val)
		newJson := fmt.Sprintf("%+v\n", record)
		utils.LogStr("updating existing record for topic ", msg.Topic(), "to", newJson)
	}

	marshalled, _ := json.Marshal(record)
	return marshalled
}

func (processor SensorProcessor) IsProcessable(topic string) bool {
	return utils.Contains(sensors, topic)
}

func (processor SensorProcessor) createNewRecord(msg MQTT.Message) model.SensorRecord {
	var data model.SensorData
	err2 := json.Unmarshal(msg.Payload(), &data)
	utils.PrintIfErr(err2)
	currentHour := float64(time.Now().Hour())
	if data.Occupancy {
		return model.BuildSensorRecord(1, 0, currentHour)
	}
	return model.BuildSensorRecord(0, 1, currentHour)
}

func (processor SensorProcessor) updateExistingRecord(msg MQTT.Message, val string) model.SensorRecord {
	var data model.SensorData
	var oldRecord model.SensorRecord
	_ = json.Unmarshal(msg.Payload(), &data)
	_ = json.Unmarshal([]byte(val), &oldRecord)

	now := float64(time.Now().Hour())
	average := utils.CalculateRollingAverage(oldRecord.AverageOccupanyTime, oldRecord.MotionCleared+oldRecord.MotionDetected+1, now)

	if data.Occupancy {
		return model.BuildSensorRecord(oldRecord.MotionDetected+1, oldRecord.MotionCleared, average)
	}
	return model.BuildSensorRecord(oldRecord.MotionDetected, oldRecord.MotionCleared+1, average)
}
