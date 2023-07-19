package lrucache

import (
	"log"
	"sync"

	"github.com/anosovs/rebrain_LRU/internal/list"
)


type LRUCache interface {
    // Добавляет новое значение с ключом в кеш (с наивысшим приоритетом), возвращает true, если все прошло успешно
    // В случае дублирования ключа вернуть false
    // В случае превышения размера - вытесняется наименее приоритетный элемент
    Add(key, value string) bool

    // Возвращает значение под ключом и флаг его наличия в кеше
    // В случае наличия в кеше элемента повышает его приоритет
    Get(key string) (value string, ok bool)

    // Удаляет элемент из кеша, в случае успеха возврашает true, в случае отсутствия элемента - false
    Remove(key string) (ok bool)
}

type cache struct {
	sync.RWMutex
	maxCount int
	cache map[string]string
	queue list.ListInterface
}

func NewLRUCache(n int) LRUCache {
	return &cache{
		maxCount: n,
		cache: make(map[string]string),
		queue: list.InitList(n),
	}
}


func (c *cache) Add(key, value string) bool {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.cache[key]; ok {
		return false
	}
	c.cache[key] = value
	
	if removedKey := c.queue.AddFront(key); removedKey!="" {
		delete(c.cache, removedKey)
	}
	// DEBUG Показывает список
	// c.queue.Traverse()
	return true
}

func (c *cache) Get(key string) (value string, ok bool) {
	c.Lock()
	defer c.Unlock()
	if record, ok := c.cache[key]; ok {
		err := c.queue.MoveToFront(key)
		if err != nil {
			log.Println(err)
		}
		return record, true
	}
	return "", false
}

func (c *cache) Remove(key string) (ok bool) {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.Get(key); ok {
		delete(c.cache, key)
		c.queue.RemoveByName(key)
		return true
	}
	return false	
}