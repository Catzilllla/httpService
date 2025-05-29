package cachemod

import (
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"ipocalc/ipocalc/internal/models"
	"log"
	"sync"
	"time"
)

func HashRequestBody(body models.JsonRequest) string {
	jsonToBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	h := fnv.New32a()
	h.Write(jsonToBytes)
	return fmt.Sprintf("%d", h.Sum32())
}

type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	CleanupInterval   time.Duration
	Citems            map[string]Item
	Counter           int
}

func NewContainer(defaultExpiration, cleanupInterval time.Duration) *Cache {
	cache := Cache{
		Citems:            map[string]Item{},
		defaultExpiration: defaultExpiration,
		CleanupInterval:   cleanupInterval,
		Counter:           0,
	}
	if cleanupInterval > 0 {
		cache.StartGC()
	}
	return &cache
}

func (c *Cache) GetAll() map[string]interface{} {
	c.RLock()
	defer c.RUnlock()

	// Создаем новый map для возвращаемых данных
	result := make(map[string]interface{})

	for key, item := range c.Citems {
		if item.Expiration == 0 || item.Expiration > time.Now().Unix() {
			result[key] = item.Value
		}
	}
	return result
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {

	var expiration int64

	// Если продолжительность жизни равна 0 - используется значение по-умолчанию
	if duration == 0 {
		duration = c.defaultExpiration
	}

	// Устанавливаем время истечения кеша
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.Lock()

	defer c.Unlock()

	c.Citems[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {

	c.RLock()

	defer c.RUnlock()

	item, found := c.Citems[key]

	// ключ не найден
	if !found {
		return nil, false
	}

	// Проверка на установку времени истечения, в противном случае он бессрочный
	if item.Expiration > 0 {

		// Если в момент запроса кеш устарел возвращаем nil
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}

	}

	return item.Value, true
}

func (c *Cache) Delete(key string) error {

	c.Lock()

	defer c.Unlock()

	if _, found := c.Citems[key]; !found {
		return errors.New("key not found")
	}

	delete(c.Citems, key)

	return nil
}

func (c *Cache) StartGC() {
	go c.GC()
}

func (c *Cache) GC() {

	for {
		// ожидаем время установленное в cleanupInterval
		<-time.After(c.CleanupInterval)

		if c.Citems == nil {
			return
		}

		// Ищем элементы с истекшим временем жизни и удаляем из хранилища
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)

		}

	}

}

// expiredKeys возвращает список "просроченных" ключей
func (c *Cache) expiredKeys() (keys []string) {

	c.RLock()

	defer c.RUnlock()

	for k, i := range c.Citems {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

// clearItems удаляет ключи из переданного списка, в нашем случае "просроченные"
func (c *Cache) clearItems(keys []string) {

	c.Lock()

	defer c.Unlock()

	for _, k := range keys {
		delete(c.Citems, k)
	}
}
