package redis

import (
	"github.com/Ho-J/base/config"
	"github.com/redis/go-redis/v9"
)

var redisM = make(map[string]*redis.Client)
var defaultDB = "default"

func NewCilent(redisConns []config.Redis) map[string]*redis.Client {
	dbM := make(map[string]*redis.Client)
	for i, conn := range redisConns {
		opt, err := redis.ParseURL(conn.Conn) // "redis://<user>:<pass>@localhost:6379/<db>"
		if err != nil {
			panic(err)
		}

		dbM[conn.DbName] = redis.NewClient(opt)

		if i == 0 {
			dbM[defaultDB] = dbM[conn.DbName]
		}
	}

	return dbM
}

func InitRedis(redisConns []config.Redis) {
	redisM = NewCilent(redisConns)
}

func GetDB(dbName ...string) *redis.Client {
	if len(dbName) == 0 {
		return redisM[defaultDB]
	}

	return redisM[dbName[0]]
}
