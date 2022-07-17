package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	opt := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	client := redis.NewClient(opt)

	if err := client.Set(ctx, "foo", "bar", 0).Err(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println(client.Get(ctx, "foo").Result())
}
