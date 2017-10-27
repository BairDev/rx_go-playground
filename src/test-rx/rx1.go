package main
import (
	"fmt"
	"errors"
	"google.golang.org/api/iterator"
	"github.com/reactivex/rxgo/observer"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/handlers"
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

	it, _ := iterable.New([]interface{}{1, 2, 3, 4, errors.New("bang"), 5})
	source := observable.From(it)
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

// smth with iterable
// input := make(chan interface{})
// iter, _ := iterable.New(input)
// source := From(iter)

// go func() {
// 	input <- "hello world"
// 	close(input)
// }()

// <-source.Subscribe(handlers.NextFunc(func(value interface{}) {
// 	// do work here
// }))