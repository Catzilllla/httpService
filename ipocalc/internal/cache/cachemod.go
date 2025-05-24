package cachemod

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"time"

	"github.com/patrickmn/go-cache"
)

// Хэширование тела запроса (для использования как ключа кэша)
func HashRequestBody(body []byte) string {
	h := fnv.New32a()
	h.Write(body)
	return fmt.Sprintf("%d", h.Sum32())
}

// Структура для хранения кэша
type CacheStore struct {
	cache    *cache.Cache
	keycache []string
}

// Создание нового кэша с указанным TTL
func NewCacheStore(defaultTTL, cleanupInterval time.Duration) *CacheStore {
	return &CacheStore{
		cache:    cache.New(defaultTTL, cleanupInterval),
		keycache: make([]string, 0),
	}
}

// Сохранение в кэш с TTL
func (c *CacheStore) Set(key string, data interface{}, ttl time.Duration) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Устанавливаем значение с определенным TTL
	c.cache.Set(key, jsonBytes, ttl)
	return nil
}

// Извлечение из кэша
func (c *CacheStore) Get(key string, dest interface{}) error {
	value, found := c.cache.Get(key)
	if !found {
		return fmt.Errorf("ключ %s не найден", key)
	}
	return json.Unmarshal(value.([]byte), dest)
}

func (c *CacheStore) GetAll() map[string]interface{} {
	allItems := make(map[string]interface{})
	for _, key := range c.keycache {
		if value, found := c.cache.Get(key); found {
			allItems[key] = value
		}
	}
	return allItems
}
