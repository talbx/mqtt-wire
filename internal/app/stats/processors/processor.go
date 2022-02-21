package processors

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Processor interface {
	Process(msg MQTT.Message, val string) []byte
	IsProcessable(topic string) bool
}
