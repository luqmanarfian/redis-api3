package main

import (
	"encoding/json"
	"fmt"

	"redis-api3/author"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "34.68.30.186:7000",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(author.Author{Name: "Harimau", Age: 3})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id345678", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("id345678").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

}
