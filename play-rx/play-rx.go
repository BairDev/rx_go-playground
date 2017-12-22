package main

import (
	"fmt"
	"rx_go/rx_basic/simple_observable"
	"rx_go/rx_basic/simple_observer"
)

func printLater() {
	defer fmt.Println("Waited for printing 'this'. That!")

	fmt.Println("This.");
}

func main() {

	printLater()

	subscription := simple_observable.GetSimpleObservable().Subscribe(simple_observer.GetSimpleObserver())
	<- subscription
}
