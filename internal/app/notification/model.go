package service

type InputData struct {
	DataPoints []DataPoint
}

type DataPoint struct {
	Message Message `json:"message"`

}

type Message struct {
	BatteryStatus int `json:"battery"`
	Topic string `json:"topic"`
}
