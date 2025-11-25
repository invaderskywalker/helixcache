package cache

import "go.uber.org/zap"

type ICache interface {
	Set(key string, val []byte) error
	Get(key string) ([]byte, bool)
	Delete(key string) error
}

// Optionally allow future logger injection via interface
type LoggerAware interface {
	SetLogger(logger *zap.Logger)
}
