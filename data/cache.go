package data

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"

	"github.com/makinori/blahaj-quest/config"
)

var cacheMutex sync.Mutex

type cacheEntry struct {
	Data   json.RawMessage `json:"data"`
	Expire time.Time       `json:"expire"`
}

type cacheMap map[string]cacheEntry

func getCacheMap() (cacheMap, error) {
	cacheBytes, err := os.ReadFile(config.CacheJSONPath)
	if err != nil {
		return cacheMap{}, nil // just init empty
	}

	var cache cacheMap

	err = json.Unmarshal(cacheBytes, &cache)
	if err != nil {
		return cacheMap{}, errors.New("failed to unmarshal cache: " + err.Error())
	}

	return cache, nil
}

func GetCache[T any](key string, value *T) error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	cache, err := getCacheMap()
	if err != nil {
		return err
	}

	entry, ok := cache[key]
	if !ok {
		return errors.New("failed to find entry")
	}

	if time.Now().After(entry.Expire) {
		return errors.New("cache entry expired")
	}

	err = json.Unmarshal(entry.Data, value)
	if err != nil {
		return errors.New("failed to cast cache: " + err.Error())
	}

	return nil
}

func SetCache(key string, value any) error {
	valueJson, err := json.Marshal(value)
	if err != nil {
		return errors.New("failed to marshal cache value: " + err.Error())
	}

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	cache, err := getCacheMap()
	if err != nil {
		return err
	}

	cache[key] = cacheEntry{
		Data:   valueJson,
		Expire: time.Now().Add(config.CacheExpireTime),
	}

	cacheData, err := json.Marshal(&cache)
	if err != nil {
		return errors.New("failed to marshal cache: " + err.Error())
	}

	err = os.WriteFile(config.CacheJSONPath, cacheData, 0644)
	if err != nil {
		return errors.New("failed to write cache: " + err.Error())
	}

	return nil
}
