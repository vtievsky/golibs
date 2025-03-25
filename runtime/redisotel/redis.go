package redisotel

import (
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

// NewUniversalClient создает redis.UniversalClient и настраивает для него метрики и трейсинг.
func NewUniversalClient(opts *redis.UniversalOptions) (redis.UniversalClient, error) {
	rdb := redis.NewUniversalClient(opts)

	return WithOtel(rdb)
}

// WithOtel оборачивает клиент redis в метрики и трейсинг.
func WithOtel(rdb redis.UniversalClient) (redis.UniversalClient, error) {
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		return nil, err
	}

	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		return nil, err
	}

	return rdb, nil
}
