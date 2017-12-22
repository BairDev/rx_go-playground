package simple_observable

import (
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/iterable"
)

func GetSimpleObservable() observable.Observable {
	rxIterable, _ := iterable.New([]interface{}{1, 2, 3, 4, 5})

	// here is room for modifications
	source := observable.From(rxIterable)
	return source.Take(3)
}

func GetIntObservable(numList iterable.Iterable) observable.Observable {
	return observable.From(numList)
}