package lock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/limeschool/gin"
	"service/consts"
	"time"
)

type lock struct {
	redis    *redis.Client
	key      string
	val      string
	duration time.Duration
}

type Lock interface {
	Acquire()
	TryAcquire() bool
	Release()
}

func NewLockWithDuration(ctx *gin.Context, key string, duration time.Duration) Lock {
	return &lock{
		redis:    ctx.Redis(consts.RedisLock),
		key:      key,
		duration: duration,
	}
}

func NewLock(ctx *gin.Context, key string) Lock {
	return &lock{
		redis:    ctx.Redis(consts.RedisLock),
		key:      key,
		duration: 30 * time.Second,
	}
}

// Acquire 获取分布式锁
func (l *lock) Acquire() {
	for {
		// 获得锁
		if res := l.redis.SetNX(context.TODO(), l.key, true, l.duration); res.Err() == nil && res.Val() {
			break
		}
	}
}

// TryAcquire 尝试获取锁，不会阻塞
func (l *lock) TryAcquire() bool {
	if res := l.redis.SetNX(context.TODO(), l.key, true, l.duration); res.Err() == nil && res.Val() {
		return true
	}
	return false
}

func (l *lock) Release() {
	l.redis.Del(context.TODO(), l.key)
}
