package main

import (
	"fmt"
	"rx_go/rx_basic/merge_observables"
	"rx_go/rx_basic/simple_observable"
	"rx_go/rx_basic/simple_observer"
	"sync"

	"github.com/reactivex/rxgo/iterable"
)

// synchronizing pattern, actually for main package/method https://stackoverflow.com/a/26928625/2092322
var waitGroup sync.WaitGroup

func printLater() {
	defer fmt.Println("Waited for printing 'this'. That!")

	fmt.Println("This.")
}

func main() {

	printLater()

	subscription := simple_observable.GetSimpleObservable().Subscribe(simple_observer.GetSimpleObserver())
	<-subscription

	rxNumIterable1, _ := iterable.New([]interface{}{1, 2, 3, 4, 5})
	rxNumIterable2, _ := iterable.New([]interface{}{6, 7, 8, 9, 10})
	rxNumIterable3, _ := iterable.New([]interface{}{11, 12, 13, 14, 15})

	observable1 := simple_observable.GetIntObservable(rxNumIterable1)
	observable2 := simple_observable.GetIntObservable(rxNumIterable2)
	observable3 := simple_observable.GetIntObservable(rxNumIterable3)

	subscriptionMerged := merge_observables.Merge(&waitGroup, observable1, observable2, observable3).Subscribe(simple_observer.GetSimpleObserver())
	<-subscriptionMerged

	waitGroup.Wait()
}
