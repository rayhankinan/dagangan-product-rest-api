package cache

import (
	"sync"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"

	cacheConfig "dagangan-product-rest-api/config/cache"
)

type LocalCache struct {
	store *persistence.InMemoryStore
	once  sync.Once
}

// Private
func (localCache *LocalCache) lazyInit() {
	localCache.once.Do(func() {
		config := cacheConfig.Config.GetMetadata()
		localCache.store = persistence.NewInMemoryStore(config.ExpirationTime)
	})
}

// Public
func (localCache *LocalCache) GetHandlerFunc(handle gin.HandlerFunc) gin.HandlerFunc {
	localCache.lazyInit()
	config := cacheConfig.Config.GetMetadata()
	return cache.CachePage(localCache.store, config.ExpirationTime, handle)
}

var Store = &LocalCache{}
