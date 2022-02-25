package notification

import (
	"github.com/gregdel/pushover"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockPushService struct {
	mock.Mock
}

func (m *mockPushService) Notify(_ *pushover.Message) *pushover.Response {
	return &pushover.Response{
		ID: "mock",
	}
}

func TestShouldNotProcess(t *testing.T) {
	deviceData := DeviceData{
		Message: Message{BatteryStatus: 57},
	}
	mockService := new(mockPushService)
	processor := StatusProcessor{mockService}
	resp := processor.ProcessStatus(deviceData)
	assert.Nil(t, resp, "no further processing since battery status too high")
}

func TestShouldProcess(t *testing.T) {
	deviceData := DeviceData{
		Message: Message{BatteryStatus: 8},
	}
	mockService := new(mockPushService)
	mockService.On("notify", deviceData).Return(&pushover.Response{
		ID: "mock",
	})

	processor := StatusProcessor{mockService}
	resp := processor.ProcessStatus(deviceData)
	assert.NotNil(t, resp, "processing should be done")
	assert.Equal(t, resp.ID, "mock")
}
