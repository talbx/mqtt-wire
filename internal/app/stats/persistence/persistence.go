package persistence

import (
	"context"
	Utils "github.com/talbx/mqtt-wire/internal/pkg/utils"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var RedisClient = redis.NewClient(&redis.Options{
	Addr:     Utils.WireConf.Redis.Host,
	Password: Utils.WireConf.Redis.Password,
	DB:       int(Utils.WireConf.Redis.Db),
})

func GetRecord(topic string) string {
	result, _ := RedisClient.Get(ctx, topic).Result()
	return result
}

func PersistRecord(data []byte, topic string) {
	err := RedisClient.Set(ctx, topic, data, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	Utils.LogStr("Sucessfully persisted record for topic ", topic)
}
