package cachemap

import "sync"

type DurableMap3 struct {
	smap  sync.Map
	store Store3
}

func NewDurableMap3(store Store3) (*DurableMap3, error) {
	o := DurableMap3{store: store}
	err := store.Load(&o.smap)
	if err != nil {
		return nil, err
	}
	return &o, err
}

func (o *DurableMap3) Get(k interface{}) (v interface{}, err error) {
	v, pres := o.smap.Load(k)
	if pres {
		return v, nil
	} else {
		return nil, nil
	}
}
