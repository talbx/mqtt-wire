package delegation

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/talbx/mqtt-wire/internal/app/stats/persistence"
	"github.com/talbx/mqtt-wire/internal/app/stats/processors"
	Utils "github.com/talbx/mqtt-wire/internal/pkg/utils"
)

var lightProcessor = processors.LightsProcessor{}
var sensorProcessor = processors.SensorProcessor{}
var radiatorProcessor = processors.RadiatorProcessor{}
var procs = [...]processors.Processor{lightProcessor, sensorProcessor, radiatorProcessor}

func Delegate(msg MQTT.Message) []byte {
	for _, processor := range procs {
		if processor.IsProcessable(msg.Topic()) {
			Utils.LogStr("processable: ", msg.Topic())
			dbRecord := persistence.GetRecord(msg.Topic())
			return processor.Process(msg, dbRecord)
		}
	}
	return nil
}
