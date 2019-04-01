package gopenid_test

import (
	"fmt"
	"github.com/tanopwan/gopenid"
	"testing"
	"time"
)

// CacheService contains business logic to the domain
type CacheService struct {
	data map[string]interface{}
}

// NewCacheService return new object
func NewCacheService() *CacheService {
	return &CacheService{
		data: make(map[string]interface{}),
	}
}

// Get key from cache
func (c *CacheService) Get(key string) interface{} {
	return c.data[key]
}

// Del key from cache
func (c *CacheService) Del(key string) {
	delete(c.data, key)
}

// Set key into cache
func (c *CacheService) Set(key string, value interface{}) {
	c.data[key] = value
}

// SetExpire key into cache with expiry date
func (c *CacheService) SetExpire(key string, value interface{}, duration time.Duration) {
	c.data[key] = value
	timer := time.NewTimer(duration)
	go func() {
		<-timer.C
		delete(c.data, key)
		fmt.Println("google_public cache is expired")
	}()
}


func TestValidateTokenForProd(t *testing.T) {
	cc := NewCacheService()
	s := gopenid.NewGoogleService(cc)
	idToken := ""
	claims, err := s.TokenInfoForProd(idToken)
	if err != nil {
		t.Errorf("failed test: %s\n", err.Error())
	}

	t.Logf("claims: %+v\n", claims)
}
