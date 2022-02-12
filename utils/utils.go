package utils

import (
	"fmt"
	"github.com/talbx/mqtt-wire/model"
	"io/ioutil"
	"log"
	"math"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"gopkg.in/yaml.v3"
)

var WireConf, _ = ReadConf()

func CalculateRollingAverage(currentAverage float64, n int, newVal float64) float64 {
	var avg = toFixed(currentAverage,2)
	avg -= toFixed(avg / float64(n),2)
	avg += toFixed(newVal / float64(n),2)
	return avg
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}

func PrintIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ReadConf() (*model.WireConfig, error) {
	filename := "config.yaml"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &model.WireConfig{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	fmt.Println(c)
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
