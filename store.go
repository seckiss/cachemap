package cachemap

import (
	"sync"

	"github.com/davecgh/go-spew/spew"
)

func p(fs string, args ...interface{}) {
	//fmt.Printf(fs+"\n", args...)
	spew.Printf(fs+"\n", args...)
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

type Store1 interface {
	Load(internalMap *sync.Map) error
	Insert(k, v interface{}) error
	Update(k, v interface{}) error
}

type Store2 interface {
	Load(internalMap *sync.Map) error
	Insert(k, v interface{}) error
}

type Store3 interface {
	Load(internalMap *sync.Map) error
}
