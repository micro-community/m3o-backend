package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type limiter struct {
	c *redis.Client
}

const (
	prefixRateLimit = "v1-service/ratelimit"
)

var (
	errLimitExceeded = fmt.Errorf("limit exceeded")
)

// incr increments the counter for this ID and compares against limitPerSec. If the incr operation pushes it over the limit
// for this second an error is returned
func (l *limiter) incr(id string, limitPerSec int64) error {
	ctx := context.Background()
	pipe := l.c.TxPipeline()
	key := fmt.Sprintf("%s:%s:%2d", prefixRateLimit, id, time.Now().Second())
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, 1*time.Second)
	pipe.Exec(ctx)
	res, err := incr.Result()
	if err != nil {
		return err
	}
	if res > limitPerSec {
		return errLimitExceeded
	}
	return nil
}
