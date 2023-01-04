package lru_cache

import (
	"testing"
)

var testData = []struct {
	key   string
	value interface{}
}{
	{
		key:   "aaa",
		value: "111",
	},
	{
		key:   "bbb",
		value: "222",
	},
	{
		key:   "ccc",
		value: "333",
	},
	{
		key:   "ddd",
		value: "444",
	},
}

// go test -run TestLruCache$
func TestLruCache(t *testing.T) {
	lrc := NewLrc(3)
	for _, item := range testData {
		lrc.Set(item.key, item.value)
	}
	bbb, _ := lrc.Get("bbb")
	if bbb != "222" {
		t.Error("assert valBb error")
	}
	aaa, _ := lrc.Get("aaa")
	if aaa != nil {
		t.Error("assert aaa error")
	}

	lenList := lrc.list.Len()
	lenMap := len(lrc.cacheMap)
	if lenList != 3 || lenMap != 3 {
		t.Error("assert len error")
	}

	lrc.Delete("bbb")
	lenList = lrc.list.Len()
	lenMap = len(lrc.cacheMap)
	if lenList != 2 || lenMap != 2 {
		t.Log(lenList, lenMap)
		t.Error("assert len error")
	}
}

func TestLruCache_repeatSet(t *testing.T) {
	lrc := NewLrc(3)
	key := "foo"
	lrc.Set(key, 1)
	val, _ := lrc.Get(key)
	if val != 1 {
		t.Errorf("got %v, want 1", val)
	}
	lrc.Set(key, 2)
	val, _ = lrc.Get(key)
	if val != 2 {
		t.Errorf("got %v, want 2", val)
	}
}

func TestLruCache_base(t *testing.T) {
	lrc := NewLrc(3)
	key := "foo"
	value := "bar"
	lrc.Set(key, value)
	val, ok := lrc.Get(key)
	if val != value || ok != true {
		t.Error("assert error")
	}
	val, ok = lrc.Get("not set")
	if val != nil || ok != false {
		t.Error("assert nonExist but exist")
	}
}

func TestLruCache_emptyGET(t *testing.T) {
	lrc := NewLrc(3)
	key := "foo"
	lrc.Get(key)
	_, ok := lrc.Get(key)
	if ok != false {
		t.Errorf("got %t, want false", ok)
		return
	}
}
