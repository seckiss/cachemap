package cachemap

import "sync"

type DurableMap2 struct {
	smap     sync.Map
	store    Store2
	resource Resource
}

func NewDurableMap2(store Store2, resource Resource) (*DurableMap2, error) {
	o := DurableMap2{store: store, resource: resource}
	err := store.Load(&o.smap)
	if err != nil {
		return nil, err
	}
	return &o, err
}

func (o *DurableMap2) Get(k interface{}) (v interface{}, err error) {
	v, pres := o.smap.Load(k)
	if pres {
		return v, nil
	} else {
		//race conditions possible - cuncurrent Resolve() possible for the same key. We are going to live with that
		v, err = o.resource.Resolve(k)
		if err != nil {
			return nil, err
		}
		err = o.store.Insert(k, v)
		if err != nil {
			return nil, err
		}
		o.smap.Store(k, v)
		return v, nil
	}
}
