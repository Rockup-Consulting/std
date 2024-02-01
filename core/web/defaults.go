package web

import (
	"net/http"
	"time"
)

const (
	DefaultReadTimeout     = time.Second * 5
	DefaultWriteTimeout    = time.Second * 10
	DefaultIdleTimeout     = time.Second * 120
	DefaultShutdownTimeout = time.Second * 20

	// DefaultStaticCache is one year is seconds.
	// https://developer.chrome.com/docs/lighthouse/performance/uses-long-cache-ttl
	DefaultStaticCache = 31536000
)

func SetServerDefaults(server *http.Server) {
	server.MaxHeaderBytes = http.DefaultMaxHeaderBytes
	server.ReadTimeout = DefaultReadTimeout
	server.WriteTimeout = DefaultWriteTimeout
	server.IdleTimeout = DefaultIdleTimeout
}
