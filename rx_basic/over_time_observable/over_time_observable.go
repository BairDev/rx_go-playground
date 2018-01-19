package over_time_observable

import (
	"time"

	"github.com/reactivex/rxgo/observable"
)

// func OverTimeObservable(waitGroup *sync.WaitGroup) observable.Observable {
func OverTimeObservable(count int) observable.Observable {
	running_channel := make(chan interface{})

	go func() { // rather separate routine for each looping
		for i := 0; i < count; i++ {
			time.Sleep(2000000 * time.Nanosecond)
			running_channel <- i
		}
		close(running_channel)
	}()

	// onError := handerls.ErrorFunc(func(item iterface{}) {
	// 	switch item := item.(type) {
	// 	case error:
	// 		handle(item)
	// 	default:
	// 		return
	// 	}
	// })

	return observable.Observable(running_channel)
}
