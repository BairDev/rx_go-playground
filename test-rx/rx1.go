package main

import (
	"fmt"
	"errors"
	"github.com/reactivex/rxgo/observer"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/handlers"
	"github.com/reactivex/rxgo/iterable"
)

func watchSimple() {
	watcher := observer.Observer{
		
		// Register a handler function for every next available item.
		NextHandler: func(item interface{}) {
			fmt.Printf("Processing: %v\n", item)
		},
		
		// Register a handler for any emitted error.
		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
		},
		
		// Register a handler when a stream is completed.
		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}
	
	// multiple value? what does that mean?
	rxIterable, _ := iterable.New([]interface{}{1, 2, 3, 'a', errors.New("bang"), 5}) // can pour into this anything
	source := observable.From(rxIterable)
	sub := source.Subscribe(watcher)
	
	// wait for the channel to emit a Subscription
	<- sub
}

func main() {
	
	score := 9
	
	onNext := handlers.NextFunc(func(item interface{}) { // http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
		// classical type assertion https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/type-assertion-and-type-switch.html
		if num, ok := item.(int); ok {
			score += num
			fmt.Printf("Incoming: %d\n", num)
		}
	})
	
	onDone := handlers.DoneFunc(func() {
		score *= 2
	})
	
	watcher := observer.New(onNext, onDone)
	
	// Create an `Observable` from a single item and subscribe to the observer.
	sub := observable.Just(1).Subscribe(watcher)
	<- sub
	
	fmt.Printf("Score has become: %d\n", score) // 20
	
	watchSimple()
}
// former import statements
// "errors"
// "google.golang.org/api/iterator"


// go env

// GOPATH="/home/mbuechs/work/go"
// GOROOT="/usr/local/go"
// PKG_CONFIG="pkg-config"