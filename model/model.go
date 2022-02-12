package model

type WireConfig struct {
	Mqttwire struct {
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
}

type RadiatorData struct {
	SystemMode       string  `json:"system_mode"`
	LocalTemperature float64 `json:"local_temperature"`
}

type RadiatorRecord struct {
	UpdatesReceived    int     `json:"updatesReceived"`
	AverageTemperature float64 `json:"averageTemperature"`
}

type LightUpdate struct {
	State string `json:"state"`
}

type LightRecord struct {
	ToggledOn  int `json:"toggledOn"`
	ToggledOff int `json:"toggledOff"`
}

type SensorData struct {
	Occupancy bool `json:"occupancy"`
}

type SensorRecord struct {
	AverageOccupanyTime float64 `json:"averageOccupancyTime"`
	MotionDetected int `json:"motionDetected"`
	MotionCleared  int `json:"motionCleared"`
}

func BuildRadiatorRecord(updatesReceived int, avgTemp float64) RadiatorRecord {
	return RadiatorRecord{
		UpdatesReceived:    updatesReceived,
		AverageTemperature: avgTemp,
	}
}

func BuildSensorRecord(motionDetected int, motionCleared int, averageOccupancyTime float64) SensorRecord {
	return SensorRecord{
		averageOccupancyTime,
		motionDetected,
		motionCleared,
	}
}

func BuildLightRecord(toggledOn int, toggledOff int) LightRecord {
	return LightRecord{
		ToggledOn:  toggledOn,
		ToggledOff: toggledOff,
	}
}
