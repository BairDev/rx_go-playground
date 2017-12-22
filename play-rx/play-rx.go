package main

import (
	"sync"
	"fmt"
	"rx_go/rx_basic/simple_observable"
	"rx_go/rx_basic/simple_observer"
	"rx_go/rx_basic/merge_observables"
	"github.com/reactivex/rxgo/iterable"
)

// synchronizing pattern, actually for main package/method https://stackoverflow.com/a/26928625/2092322
var waitGroup sync.WaitGroup

func printLater() {
	defer fmt.Println("Waited for printing 'this'. That!")

	fmt.Println("This.");
}

func main() {

	printLater()

	subscription := simple_observable.GetSimpleObservable().Subscribe(simple_observer.GetSimpleObserver())
	<- subscription

	rxNumIterable1, _ := iterable.New([]interface{}{1, 2, 3, 4, 5})
	rxNumIterable2, _ := iterable.New([]interface{}{6, 7, 8, 9, 10})
	observable1 := simple_observable.GetIntObservable(rxNumIterable1)
	observable2 := simple_observable.GetIntObservable(rxNumIterable2)

	subscriptionMerged := merge_observables.Merge(observable1, observable2, &waitGroup).Subscribe(simple_observer.GetSimpleObserver())
	<- subscriptionMerged

	waitGroup.Wait()
}
