package gin

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore"
)

func SetUp(addr string) (err error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	// New default RedisStore
	if store, err = redisstore.NewRedisStore(client); err != nil {
		return err
	}
	store.KeyPrefix("session_")
	store.Options(sessions.Options{
		MaxAge:   86400 * 60,
		Secure:   false,
		HttpOnly: true,
	})
	return nil
}