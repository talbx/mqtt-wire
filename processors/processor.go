package processors

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const prefix = "zigbee2mqtt"

type Processor interface {
	Process(msg MQTT.Message, val string) []byte
	IsProcessable(topic string) bool
}
