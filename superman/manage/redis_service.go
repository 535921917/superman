package manage

import (
	"context"
	"time"
)

func CmdsExec(cmdName, key string, value interface{}, args ...interface{}) {
	conn := pool.Get()

	_, _ = conn.Do(cmdName, "shuai", "1243")

}

func Set(ctx context.Context, key string) {
	RedisClient.Set(ctx, key, "test", time.Minute)
	_, _ = RedisClient.Get(ctx, key).Result()
}
