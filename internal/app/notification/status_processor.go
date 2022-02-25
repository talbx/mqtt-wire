package notification

import (
	"fmt"
	"github.com/gregdel/pushover"
	"log"
)

const critical = 10

type StatusProcessor struct {
	NotificationService NotificationService
}

func (processor StatusProcessor) ProcessStatus(deviceData DeviceData) *pushover.Response {
	if deviceData.Message.BatteryStatus <= critical {
		msg := fmt.Sprintf("%s battery status is only %d percent", deviceData.Topic, deviceData.Message.BatteryStatus)
		message := &pushover.Message{
			Message:  msg,
			Title:    "mqtt-wire",
			Priority: pushover.PriorityHigh,
			Sound:    pushover.SoundCosmic,
		}
		return processor.NotificationService.Notify(message)
	}
	log.Printf("No need for processing %s, since battery status is okay", deviceData.Topic)
	return nil
}
