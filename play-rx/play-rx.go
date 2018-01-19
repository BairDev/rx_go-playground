package main

import (
	"fmt"
	"rx_go/rx_basic/merge_observables"
	"rx_go/rx_basic/mimic_delay_observable"
	"rx_go/rx_basic/simple_observable"
	"rx_go/rx_basic/simple_observer"
	"sync"

	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
)

// synchronizing pattern, actually for main package/method https://stackoverflow.com/a/26928625/2092322
var waitGroup sync.WaitGroup

func printLater() {
	defer fmt.Println("Waited for printing 'this'. That!")

	fmt.Println("This.")
}

func multiplierMap(number interface{}) interface{} {
	numericalVal, ok := number.(int)
	if ok {
		return (numericalVal * 3)
	} else {
		fmt.Println("No number in multiplierMap.")
	}

	return 0
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

	observables := []observable.Observable{observable1, observable2, observable3}

	subscriptionMerged := merge_observables.Merge(&waitGroup, observables).Subscribe(simple_observer.GetSimpleObserver())
	<-subscriptionMerged

	// this is a blocking wait
	waitGroup.Wait()

	rxNumIterable4, _ := iterable.New([]interface{}{10, 20, 30, 20, 10})
	mappingObservable := simple_observable.GetIntObservable(rxNumIterable4).Map(multiplierMap)
	subscriptionMapped := mappingObservable.Subscribe(simple_observer.GetSimpleObserver())
	<-subscriptionMapped

	delayingObservable := mimic_delay_observable.GetWithDelayingEmitter()
	subscriptionWait := delayingObservable.Subscribe(simple_observer.GetSimpleObserver())
	<-subscriptionWait
}
