package cache

import (
	"scrape/domain"
	"sync"

	"scrape/domain/errorCode"
)

type Cache struct {
	mutex sync.RWMutex
	List  map[string]map[int]*domain.House
}

func NewCache(names ...string) (c *Cache) {
	c = &Cache{
		mutex: sync.RWMutex{},
		List:  make(map[string](map[int]*domain.House)),
	}
	for _, eachName := range names {
		c.List[eachName] = make(map[int]*domain.House)
	}
	return
}

func (c *Cache) Add(source string, req *domain.House) (code int, data interface{}, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.List[source][req.HouseId] = req
	return
}

func (c *Cache) GetAll() (code int, data map[string]map[int]*domain.House, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	data = c.List
	return
}

func (c *Cache) IsDataExist(source string, req *domain.House) (code int, data interface{}, err error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	tmpData, isExist := c.List[source][req.HouseId]
	if isExist {
		if req.Price != tmpData.Price {
			code = errorCode.DBNoData
			return
		}
		if req.IsDownPrice != tmpData.IsDownPrice {
			code = errorCode.DBNoData
			return
		}
		return
	}
	code = errorCode.DBNoData
	return
}
