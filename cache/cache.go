package cache

import "github.com/bradfitz/gomemcache/memcache"

const (
	mcServer = "127.0.0.1:11211"
)

var mc = memcache.New(mcServer)

// Wrapper get the Item stored in Memcache based in the key
func Get(key string) (item *memcache.Item, err error){
	return mc.Get(key)
}
// Wrapper. Stored the item in Memcache.
func Set(key string, value string)(err error){
	item := &memcache.Item{Key: key, Value: []byte(value)}
	return mc.Set(item)
}