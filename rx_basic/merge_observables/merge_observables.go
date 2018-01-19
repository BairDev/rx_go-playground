package merge_observables

import (
	"fmt"
	"sync"

	"github.com/reactivex/rxgo/handlers"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func Merge(waitGroup *sync.WaitGroup, observables []observable.Observable) observable.Observable {
	merged_channel := make(chan interface{}) // not buffered, synchronous: https://stackoverflow.com/a/5983572/2092322

	observables_number := len(observables)
	observables_closed := 0

	onNext := handlers.NextFunc(func(item interface{}) {
		switch item := item.(type) {
		case error:
			fmt.Println("onNext has encounted an error!")
		default:
			merged_channel <- item
		}
	})

	// onError := handerls.ErrorFunc(func(item iterface{}) {
	// 	switch item := item.(type) {
	// 	case error:
	// 		handle(item)
	// 	default:
	// 		return
	// 	}
	// })

	onDone := handlers.DoneFunc(func() {
		defer waitGroup.Done() // defer or not ?

		observables_closed = observables_closed + 1
		if observables_closed == observables_number {
			close(merged_channel)
		}
	})

	waitGroup.Add(observables_number)
	for i := 0; i < observables_number; i++ {
		observables[i].Subscribe(observer.New(onNext, onDone))
	}

	return observable.Observable(merged_channel)
}
