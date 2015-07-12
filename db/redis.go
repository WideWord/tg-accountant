package db

import(
	"../config"
	"gopkg.in/redis.v3"
)

var client *redis.Client

func Connect() error {
	client = redis.NewClient(&redis.Options{
        Addr:     config.Get().Redis.Host,
        Password: config.Get().Redis.Password,
        DB:       0,  // use default DB
    })

    _, err := client.Ping().Result()

    return err
}

func Client() *redis.Client {
	return client
}
