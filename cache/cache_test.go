package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
	"testing"
)

func TestErrCacheMiss(t *testing.T) {
	item, err := Get("undefined")

	if err != memcache.ErrCacheMiss {
		t.Errorf("Unexpected messager error: |%v|. Expected: |%v|", err.Error(), memcache.ErrCacheMiss)
	}
	if err == nil {
		t.Errorf("Error. Cache returns not error when key really doesnt exists %v", err.Error())
	}
	if item != nil {
		t.Error("Error. Item unexisting is not Nil")
	}
}


func TestSetItem(t *testing.T){
	key, value := "Hello", "World"

	err := Set(key, value); if err!=nil{
		t.Fatalf("Error Storing test element in memcache: %v", err.Error())
	}

	//check item stored...
	item, _ := Get(key)
	if string(item.Value) != value{
		t.Errorf("The Item stored doesn't the expected: %v , expected: %v", item.Value, value)
	}
}