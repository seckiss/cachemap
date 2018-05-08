# cachemap
Cache Map that synchronizes content in Golang


### DurableMap1

DurableMap1 is caching Put(k,v) requests. It preloads the map from Store and checks if Put(k,v) needs a roundtrip to the Store to Insert or Update the record for the given key.

It applies the following logic for Put(k,v) for the given key:

- if entry exists and its value is the same in cache, do nothing
- if entry exists and its value is different do Store.Update and updates cache
- if entry does not exist do Store.Insert and put to cache



### DurableMap2

DurableMap2 is caching Get(k) requests. It preloads the map from Store and checks if Get(k) is in cache (and implied in Store). If not it calls resolve() to get value v. It is assumed that resolve() is an expensive operation.

It applies the following logic for Get(k) for the given key:

- if entry exists return cached value
- if entry does not exist call resolve() go fetch value v and then call Store.Insert() and also put to cache
