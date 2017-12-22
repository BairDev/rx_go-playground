package simple_observer

import (
	"fmt"
	"github.com/reactivex/rxgo/observer"
)

// next, error, done handler, which just print
func GetSimpleObserver() observer.Observer {
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
	
	return watcher
}