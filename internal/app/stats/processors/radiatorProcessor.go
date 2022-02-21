package processors

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	Model "github.com/talbx/mqtt-wire/internal/app/stats/model"
	Utils "github.com/talbx/mqtt-wire/internal/pkg/utils"
)

type RadiatorProcessor struct{}

func (processor RadiatorProcessor) Process(msg MQTT.Message, val string) []byte {
	var data Model.RadiatorData
	err := json.Unmarshal(msg.Payload(), &data)
	Utils.PrintIfErr(err)
	var record Model.RadiatorRecord

	if val == "" {
		record = processor.createNewRecord(msg)
		Utils.LogStr("Creating new record for topic", msg.Topic())
	} else {
		record = processor.updateExistingRecord(msg, val)
		newJson := fmt.Sprintf("%+v\n", record)
		Utils.LogStr("updating existing record for topic ", msg.Topic(), "to", newJson)
	}

	marshalled, _ := json.Marshal(record)
	return marshalled
}

func (processor RadiatorProcessor) createNewRecord(msg MQTT.Message) Model.RadiatorRecord {
	var data Model.RadiatorData
	err2 := json.Unmarshal(msg.Payload(), &data)
	Utils.PrintIfErr(err2)
	return Model.BuildRadiatorRecord(1, data.LocalTemperature)
}

func (processor RadiatorProcessor) updateExistingRecord(msg MQTT.Message, val string) Model.RadiatorRecord {
	var data Model.RadiatorData
	var oldRecord Model.RadiatorRecord
	_ = json.Unmarshal(msg.Payload(), &data)
	_ = json.Unmarshal([]byte(val), &oldRecord)

	avg := Utils.CalculateRollingAverage(oldRecord.AverageTemperature, oldRecord.UpdatesReceived, data.LocalTemperature)
	return Model.BuildRadiatorRecord(oldRecord.UpdatesReceived+1, avg)
}

/*
{"away_mode":"OFF","battery":88,"child_lock":"LOCK","current_heating_setpoint":12,"linkquality":45,"local_temperature":19.8,"position":null,"preset":"manual","running_state":null,"system_mode":"heat","valve_detection":"ON","window_detection":"OFF"}
{"away_mode":"ON","battery":88,"child_lock":"LOCK","current_heating_setpoint":12,"linkquality":54,"local_temperature":19.8,"position":null,"preset":"away","running_state":null,"system_mode":"off","valve_detection":"ON","window_detection":"OFF"}
{"away_mode":"OFF","battery":88,"child_lock":"LOCK","current_heating_setpoint":21,"linkquality":99,"local_temperature":18.7,"position":null,"preset":"manual","running_state":null,"system_mode":"heat","valve_detection":"ON","window_detection":"OFF"}
*/
func (processor RadiatorProcessor) IsProcessable(topic string) bool {
	return Utils.Contains(Utils.WireConf.Units.Radiators, topic)
}
