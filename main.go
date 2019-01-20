package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

var client *redis.Client

func handler(w http.ResponseWriter, r *http.Request) {
	inc, err := client.Incr("counter").Result()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Fprintf(w, "Hi there, counter: %d!", inc)
}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
