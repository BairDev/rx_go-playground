package main

// DOES NOT COMPILE!

// go build -o /home/mbuechs/work/tryit/rx_go/bin/rx1 rx1.go
// rx1.go:5:2: cannot find package "google.golang.org/api/iterator" in any of:
//         /usr/local/go/src/google.golang.org/api/iterator (from $GOROOT)
//         /home/mbuechs/work/go/src/google.golang.org/api/iterator (from $GOPATH)

// go env
// GOARCH="amd64"                                                                                                                                       
// GOBIN="/home/mbuechs/work/go/bin"                                                                                                                    
// GOEXE=""
// GOHOSTARCH="amd64"
// GOHOSTOS="linux"
// GOOS="linux"
// GOPATH="/home/mbuechs/work/go"
// GORACE=""
// GOROOT="/usr/local/go"
// GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
// GCCGO="gccgo"
// CC="gcc"
// GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build327165865=/tmp/go-build -gno-record-gcc-switches"
// CXX="g++"
// CGO_ENABLED="1"
// CGO_CFLAGS="-g -O2"
// CGO_CPPFLAGS=""
// CGO_CXXFLAGS="-g -O2"
// CGO_FFLAGS="-g -O2"
// CGO_LDFLAGS="-g -O2"
// PKG_CONFIG="pkg-config"

// echo $PATH
// /home/mbuechs/.nvm/versions/node/v6.10.0/bin:/home/mbuechs/bin:/home/mbuechs/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin

// https://github.com/golang/go/wiki/GOPATH
// https://stackoverflow.com/questions/21001387/how-do-i-set-the-gopath-environment-variable-on-ubuntu-what-file-must-i-edit
// https://stackoverflow.com/questions/24306183/can-someone-explain-why-gopath-is-convenient-and-how-it-should-be-used-in-genera?noredirect=1&lq=1

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