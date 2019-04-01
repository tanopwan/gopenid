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
	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImNmMDIyYTQ5ZTk3ODYxNDhhZDBlMzc5Y2M4NTQ4NDRlMzZjM2VkYzEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiOTUzMDU2MzMzNzQxLWp2azQwMnVjZ3RoZGo1dmxiYW10dGFvcHBnbWtiZzNpLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiOTUzMDU2MzMzNzQxLWp2azQwMnVjZ3RoZGo1dmxiYW10dGFvcHBnbWtiZzNpLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTEzNzMwNDM4NTIyODExNzMyMTI5IiwiZW1haWwiOiJsM2xhY2tjYXQuZ0BnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6InVQTGNha1lRY1N4Q2sxalRPVFdQMmciLCJuYW1lIjoiU3VwaGFraWF0IExvaGFzYW1tYWt1bCIsInBpY3R1cmUiOiJodHRwczovL2xoNi5nb29nbGV1c2VyY29udGVudC5jb20vLV84OUJvb2IzQlM4L0FBQUFBQUFBQUFJL0FBQUFBQUFBQUZvL0t0SDBYSDhTWkVFL3M5Ni1jL3Bob3RvLmpwZyIsImdpdmVuX25hbWUiOiJTdXBoYWtpYXQiLCJmYW1pbHlfbmFtZSI6IkxvaGFzYW1tYWt1bCIsImxvY2FsZSI6ImVuIiwiaWF0IjoxNTUxNjY5MTUxLCJleHAiOjE1NTE2NzI3NTEsImp0aSI6ImVkNDUyNDc0YjVmYjU5MDM4MDg2MzNjYTdmY2I0YWFmZTYwZjJhZDEifQ.DWNw3AnL1dUD_DFNrBW6tFsq3gHHprx5g8qsCMiHlb59FWsj6Qso_YcNbFShDkR7Ls1yhE94LI3o20O-fOLX_xVyrxx5CCyGVX5nrd9bHyHYsX-_U2WsLypqRp-ZQ4RWMfJdlAhLFvtPOvf5ivKTZtEc85mjfx5R0rYdxlHhXe2OxdKp0P1erX2IYG0zGrMoKwFzqx47YTjHbRugnmpTjTDwWuSTY3EcWbeBXa0r7okg7wTnM02oy0B4An6aOWB2Unc0RhbtxottXjQYM-HS6IiqZqJXHuCi9nf7Isd9aQihN1d6Sr4Lc4hkkFMuyzTfQp8j75I6eYm5iDGM_YHBVg"
	claims, err := s.TokenInfoForProd(idToken)
	if err != nil {
		t.Errorf("failed test: %s\n", err.Error())
	}

	t.Logf("claims: %+v\n", claims)
}
