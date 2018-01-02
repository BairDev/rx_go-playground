package merge_observables

import (
	"fmt"
	"sync"
	"github.com/reactivex/rxgo/observer"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/handlers"
)

func Merge(observable1 observable.Observable, observable2 observable.Observable, waitGroup *sync.WaitGroup) observable.Observable {
	merged_channel := make(chan interface{}) // not buffered, synchronous: https://stackoverflow.com/a/5983572/2092322

	observables_number := 2
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
		fmt.Println("One subscription closed.")
		
		observables_closed = observables_closed + 1
		if observables_closed == observables_number {
			fmt.Println("Closing merged channel.")
			close(merged_channel)
		}
	})

	waitGroup.Add(2)
	observable1.Subscribe(observer.New(onNext, onDone))
	observable2.Subscribe(observer.New(onNext, onDone))

	// go func() {
	// 	// waitGroup.Add(1)
	// 	observable1.Subscribe(observer.New(onNext, onDone))
	// 	observable2.Subscribe(observer.New(onNext, onDone))
    //     close(merged_channel)
    // }()

	return observable.Observable(merged_channel)
}