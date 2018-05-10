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

func (o *DurableMap3) GetInt(k interface{}) (v int, err error) {
	temp, err := o.Get(k)
	return temp.(int), err
}

func (o *DurableMap3) GetInt64(k interface{}) (v int64, err error) {
	temp, err := o.Get(k)
	return temp.(int64), err
}

func (o *DurableMap3) GetString(k interface{}) (v string, err error) {
	temp, err := o.Get(k)
	return temp.(string), err
}
