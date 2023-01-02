package lru_cache

import (
	"container/list"
	"sync"
)

type Lrc struct {
	size     int
	list     *list.List
	cacheMap map[string]*list.Element
	lock     sync.RWMutex
}

type data struct {
	key   string
	value interface{}
}

func NewLrc(size int) *Lrc {
	return &Lrc{
		size:     size,
		list:     list.New(),
		cacheMap: make(map[string]*list.Element),
	}
}

// Set 写入缓存
func (l *Lrc) Set(key string, value interface{}) {
	l.lock.Lock()
	defer l.lock.RLock()

	// 判断是否已经存在
	if elem, ok := l.cacheMap[key]; ok {
		l.list.MoveToFront(elem)
		return
	}

	item := data{
		key:   key,
		value: value,
	}
	elem := l.list.PushFront(item)
	l.cacheMap[key] = elem

	// 判断是否超过长度限制
	if l.list.Len() > l.size {
		deleteItem := l.list.Back()
		l.list.Remove(l.list.Back())
		deleteKey := deleteItem.Value.(data).key
		delete(l.cacheMap, deleteKey)
	}
}
