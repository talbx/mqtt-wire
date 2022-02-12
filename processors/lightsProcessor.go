package processors

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/talbx/mqtt-wire/model"
	"github.com/talbx/mqtt-wire/utils"
)


type LightsProcessor struct{}

/**
{"brightness":251,"color_mode":"color_temp","color_temp":438,"linkquality":96,"power_on_behavior":null,"state":"ON","update":{"state":"available"},"update_available":true}
*/
var lights = []string{
	prefix + "/Tradfri Tisch-Leuchte 1",
	prefix + "/Tradfri Tisch-Leuchte 2",
	prefix + "/Tradfri Deckenleuchte 1",
	prefix + "/Innr-265-1",
	prefix + "/Innr-265-2",
	prefix + "/Innr-265-3",
	prefix + "/Innr-265-4",
	prefix + "/Innr-265-5",
}

func (processor LightsProcessor) Process(msg MQTT.Message, val string) []byte {
	var record model.LightRecord
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

func (processor LightsProcessor) IsProcessable(topic string) bool{
	return utils.Contains(lights, topic)
}

func (processor LightsProcessor) createNewRecord(msg MQTT.Message) model.LightRecord {
	var data model.LightUpdate
	err2 := json.Unmarshal(msg.Payload(), &data)
	utils.PrintIfErr(err2)
	return processor.createLightRecord(data.State, 0,0)
}

func (processor LightsProcessor) createLightRecord(state string, toggleOn int, toggleOff int) model.LightRecord {
	var record model.LightRecord
	if state == "ON" {
		record = model.LightRecord{
			ToggledOn:  toggleOn + 1,
			ToggledOff: toggleOff,
		}
	} else {
		record = model.LightRecord{
			ToggledOn:  toggleOn,
			ToggledOff: toggleOff + 1,
		}
	}
	return record
}

func (processor LightsProcessor) updateExistingRecord(msg MQTT.Message, val string) model.LightRecord {
	var data model.LightUpdate
	var oldRecord model.LightRecord
	_ = json.Unmarshal(msg.Payload(), &data)
	_ = json.Unmarshal([]byte(val), &oldRecord)
	return processor.createLightRecord(data.State, oldRecord.ToggledOn, oldRecord.ToggledOff)
}
