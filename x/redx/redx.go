package redx

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

type Conf struct {
	Addrs    []string
	UserName string
	Password string
}

var (
	DevClusterConf = Conf{
		Addrs: []string{
			REDIS_HOST_ADDRESS + ":" + REDIS_PORT_NUMBERS[0],
			REDIS_HOST_ADDRESS + ":" + REDIS_PORT_NUMBERS[1],
			REDIS_HOST_ADDRESS + ":" + REDIS_PORT_NUMBERS[2],
			REDIS_HOST_ADDRESS + ":" + REDIS_PORT_NUMBERS[3],
			REDIS_HOST_ADDRESS + ":" + REDIS_PORT_NUMBERS[4],
			REDIS_HOST_ADDRESS + ":" + REDIS_PORT_NUMBERS[5],
		},
		UserName: "",
		Password: "",
	}

	REDIS_PORT_NUMBERS = []string{"6379", "6380", "6381", "6382", "6383", "6384"}
	REDIS_HOST_ADDRESS = "127.0.0.1"
)

func NewClusterDriver(conf Conf) *redis.ClusterClient {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    conf.Addrs,
		Username: conf.UserName,
		Password: conf.Password,
	})

	return rdb
}

func NewDriver(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return rdb
}

// ClearDB is a test utility that clears all data from the current instance
func ClearDB(t testing.TB, db redis.UniversalClient) {
	db.FlushDB(context.Background())
}
