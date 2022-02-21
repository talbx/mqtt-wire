package main

import (
	"github.com/talbx/mqtt-wire/internal/app/stats/delegation"
	"github.com/talbx/mqtt-wire/internal/app/stats/persistence"
	"github.com/talbx/mqtt-wire/internal/pkg/utils"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts := createMqttOpts()
	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(utils.WireConf.Mosquitto.Topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		utils.LogStr("Connected to mosquitto instance on", utils.WireConf.Mosquitto.Broker)
	}
	<-c
}

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	processedMessage := delegation.Delegate(msg)
	if processedMessage != nil {
		persistence.PersistRecord(processedMessage, msg.Topic())
	}
}

func createMqttOpts() *MQTT.ClientOptions {
	opts := MQTT.NewClientOptions().AddBroker(utils.WireConf.Mosquitto.Broker)
	opts.SetClientID(utils.WireConf.Mosquitto.ClientId)
	opts.SetUsername(utils.WireConf.Mosquitto.Username)
	opts.SetPassword(utils.WireConf.Mosquitto.Password)
	opts.SetDefaultPublishHandler(f)
	return opts
}
