package cachemap

import "sync"

type DurableMap2 struct {
	smap  sync.Map
	store Store2
}

func NewDurableMap2(store Store2) (*DurableMap2, error) {
	o := DurableMap2{store: store}
	err := store.Load(&o.smap)
	if err != nil {
		return nil, err
	}
	return &o, err
}

func (o *DurableMap2) Get(k interface{}, resolve func() (interface{}, error)) (v interface{}, err error) {
	v, pres := o.smap.Load(k)
	if pres {
		return v, nil
	} else {
		//race conditions possible - cuncurrent Resolve() possible for the same key. We are going to live with that
		v, err = resolve()
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

func (o *DurableMap2) GetInt(k interface{}, resolve func() (interface{}, error)) (v int, err error) {
	temp, err := o.Get(k, resolve)
	return temp.(int), err
}

func (o *DurableMap2) GetInt64(k interface{}, resolve func() (interface{}, error)) (v int64, err error) {
	temp, err := o.Get(k, resolve)
	return temp.(int64), err
}

func (o *DurableMap2) GetString(k interface{}, resolve func() (interface{}, error)) (v string, err error) {
	temp, err := o.Get(k, resolve)
	return temp.(string), err
}
