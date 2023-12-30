package dao

import (
    "github.com/jinzhu/gorm"
    "github.com/go-redis/redis/v8"
	"github.com/Asong6824/douyin-user-service/global"
)

type Dao struct {
    engine *gorm.DB          // 持久化数据库
    cache  *redis.Client     // 缓存数据库
}

// NewDao 创建一个新的 Dao 实例
func NewDao(engine *gorm.DB, cache *redis.Client) *Dao {
    return &Dao{
        engine: global.DBEngine,
        cache:  global.Cache,
    }
}