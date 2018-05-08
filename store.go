package cachemap

import "sync"

type Store1 interface {
	Load(internalMap *sync.Map) error
	Insert(k, v interface{}) error
	Update(k, v interface{}) error
}

type Store2 interface {
	Load(internalMap *sync.Map) error
	Insert(k, v interface{}) error
}

type Resource interface {
	Resolve(k interface{}) (v interface{}, err error)
}
