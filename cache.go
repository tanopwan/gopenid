package gopenid

import (
	"time"
)

// Cache for key-value
type Cache interface {
	Get(key string) interface{}
	Del(key string)
	Set(key string, value interface{})
	SetExpire(key string, value interface{}, duration time.Duration)
}
