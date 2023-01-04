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
	defer l.lock.Unlock()

	// 判断是否已经存在
	if elem, ok := l.cacheMap[key]; ok {
		l.list.MoveToFront(elem)
		elem.Value.(*data).value = value
		return
	}

	item := &data{
		key:   key,
		value: value,
	}
	elem := l.list.PushFront(item)
	l.cacheMap[key] = elem

	// 判断是否超过长度限制
	if l.list.Len() > l.size && l.size > 0 {
		deleteItem := l.list.Back()
		l.list.Remove(deleteItem)
		deleteKey := deleteItem.Value.(*data).key
		delete(l.cacheMap, deleteKey)
	}
}

// Get 读取缓存
func (l *Lrc) Get(key string) (interface{}, bool) {
	l.lock.RLock()
	value, ok := l.cacheMap[key]
	l.lock.RUnlock()

	if ok {
		l.lock.Lock()
		l.list.MoveToFront(value)
		l.lock.Unlock()
		return value.Value.(*data).value, true
	} else {
		return nil, false
	}
}

// Delete 缓存
func (l *Lrc) Delete(key string) {
	l.lock.RLock()
	value, ok := l.cacheMap[key]
	l.lock.RUnlock()

	if ok {
		l.lock.Lock()
		l.list.Remove(value)
		key := value.Value.(*data).key
		delete(l.cacheMap, key)
		l.lock.Unlock()
	}
}
