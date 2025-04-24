package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func main() {

	var redisClient *redis.Client

	redisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	fmt.Println("Connected to Redis")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HELLO WORLD"))
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on :8080")
}
