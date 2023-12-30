package cache

import (
	"github.com/Asong6824/douyin-user-service/package/setting"
	"context"

	"github.com/go-redis/redis/v8"
)

func NewCache(databaseSetting *setting.CacheSetting) (*redis.Client, error) {
	Redis := redis.NewClient(&redis.Options{
		Addr:        databaseSetting.Host,
		Password:    databaseSetting.Password,
		DB:          databaseSetting.DB,
		IdleTimeout: databaseSetting.IdleTimeout,
	})
	if _, err := Redis.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return Redis, nil
}