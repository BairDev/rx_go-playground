package mimic_delay_observable

import (
	"time"

	"github.com/reactivex/rxgo/observable"
)

func emitterTwoSeconds() interface{} {

	time.Sleep(2 * time.Second)
	return 1
}

func emitterOneSecond() interface{} {

	time.Sleep(time.Second)
	return 2
}

func GetWithDelayingEmitter() observable.Observable {

	return observable.Start(emitterTwoSeconds, emitterOneSecond)
}
