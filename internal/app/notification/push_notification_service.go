package notification

import (
	"github.com/gregdel/pushover"
	"github.com/talbx/mqtt-wire/internal/pkg/utils"
	"log"
)

type NotificationService interface {
	Notify(msg *pushover.Message) *pushover.Response
}

type PushNotificationService struct {}

func (notificationService PushNotificationService) Notify(msg *pushover.Message) *pushover.Response {
	app := pushover.New(utils.WireConf.Pushover.Apitoken)
	recipient := pushover.NewRecipient(utils.WireConf.Pushover.Usertoken)
	response, err := app.SendMessage(msg, recipient)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Sent out pushover message %s", msg)
	return response
}
