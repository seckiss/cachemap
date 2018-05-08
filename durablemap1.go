package cachemap

import "sync"

type DurableMap1 struct {
	smap  sync.Map
	store Store1
}

func NewDurableMap1(store Store1) (*DurableMap1, error) {
	o := DurableMap1{store: store}
	err := store.Load(&o.smap)
	if err != nil {
		return nil, err
	}
	return &o, err
}

func (o *DurableMap1) Put(k, v interface{}) (err error) {
	actual, loaded := o.smap.LoadOrStore(k, v)
	if loaded { // it was already in the map
		if actual != v {
			//overwrite internal map with the new value
			o.smap.Store(k, v)
			//update external store
			err = o.store.Update(k, v)
			if err != nil {
				return err
			}
		} else {
			// do nothing, key already present in the map with the same value
		}
	} else {
		err = o.store.Insert(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
