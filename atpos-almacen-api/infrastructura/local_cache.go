package infrastructura

import (
	"atpos-almacen-api/dominio"
	"sync"
	"time"
)

type cacheItem[V dominio.Entity] struct {
	data                V
	expirationTimeStamp int64
}

type LocalCache[K comparable, V dominio.Entity] struct {
	timer chan struct{}
	wg    sync.WaitGroup
	mu    sync.Mutex
	data  map[K]cacheItem[V]
}

func NewLocalCache[K comparable, V dominio.Entity](cleanInterval time.Duration) *LocalCache[K, V] {
	cache := &LocalCache[K, V]{
		data:  make(map[K]cacheItem[V]),
		timer: make(chan struct{}),
	}
	cache.wg.Add(1)
	go func(cleanInterval time.Duration) {
		defer cache.wg.Done()
		cache.cleanLoop(cleanInterval)
	}(cleanInterval)
	return cache
}

func (lc *LocalCache[K, V]) cleanLoop(cleanInterval time.Duration) {
	t := time.NewTicker(cleanInterval)
	defer t.Stop()

	for {
		select {
		case <-lc.timer:
			return
		case <-t.C:
			lc.mu.Lock()
			for uid, cu := range lc.data {
				if cu.expirationTimeStamp <= time.Now().Unix() {
					delete(lc.data, uid)
				}
			}
			lc.mu.Unlock()
		}
	}
}

func (lc *LocalCache[K, V]) StopCleanLoop() {
	close(lc.timer)
	lc.wg.Wait()
}

func (lc *LocalCache[K, V]) Update(key K, data V, expirationTimeStamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.data[key] = cacheItem[V]{
		data:                data,
		expirationTimeStamp: expirationTimeStamp,
	}

}

func (lc *LocalCache[K, V]) Read(key K) (V, bool) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	d, ok := lc.data[key]
	return d.data, ok
}

func (lc *LocalCache[K, V]) Delete(key K) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	delete(lc.data, key)
}
