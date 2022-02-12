package main

import (
	"github.com/talbx/mqtt-wire/delegation"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/talbx/mqtt-wire/persistence"
	"github.com/talbx/mqtt-wire/utils"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	processedMessage := delegation.Delegate(msg)
	if processedMessage != nil {
		persistence.PersistRecord(processedMessage, msg.Topic())
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts := MQTT.NewClientOptions().AddBroker(utils.WireConf.Mqttwire.Mosquitto.Broker)
	opts.SetClientID(utils.WireConf.Mqttwire.Mosquitto.ClientId)
	opts.SetUsername(utils.WireConf.Mqttwire.Mosquitto.Username)
	opts.SetPassword(utils.WireConf.Mqttwire.Mosquitto.Password)
	opts.SetDefaultPublishHandler(f)
	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(utils.WireConf.Mqttwire.Mosquitto.Topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		utils.LogStr("Connected to mosquitto instance on", utils.WireConf.Mqttwire.Mosquitto.Broker)
	}
	<-c
}
