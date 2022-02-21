package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"gopkg.in/yaml.v3"
)

type WireConfig struct {
	Mosquitto struct {
		Broker   string
		ClientId string
		Username string
		Password string
		Topic    string
	}
	Redis struct {
		Host     string
		Password string
		Db       int8
	}
	Units struct {
		Lights    []string
		Radiators []string
		Sensors   []string
	}
}

var WireConf, _ = ReadConf()

func CalculateRollingAverage(currentAverage float64, n int, newVal float64) float64 {
	var avg = toFixed(currentAverage, 2)
	avg -= toFixed(avg/float64(n), 2)
	avg += toFixed(newVal/float64(n), 2)
	return avg
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func PrintIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ReadConf() (*WireConfig, error) {
	filename := "config.yaml"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &WireConfig{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return c, nil
}

func LogMsg(msg MQTT.Message) {
	fmt.Printf("[mqtt-wire]: %s\n", msg.Payload())
}

func LogStr(msg ...string) {
	fmt.Printf("[mqtt-wire]: %s\n", msg)
}

func Contains(list []string, element string) bool {
	for _, a := range list {
		if a == element {
			return true
		}
	}
	return false
}
