package redis

import "time"

type Redis interface {
	Get(key string) (string, error)
	Set(key string, value string, duration time.Duration)
}
