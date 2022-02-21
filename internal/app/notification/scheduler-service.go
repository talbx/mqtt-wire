package service

import (
	"github.com/talbx/mqtt-wire/internal/pkg/utils"
	"reflect"
	"runtime"
	"time"
)

type SchedulingFn func()

type SchedulerService struct {
}

func (service SchedulerService) Execute(fn SchedulingFn) {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	ticker := time.NewTicker(60 * time.Second)
	for range ticker.C {
		utils.LogStr("Executing scheduled function ", name)
		fn()
	}
}
