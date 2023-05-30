package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // Set password if your Redis server requires authentication
		DB:       0,                // Use default database
	})

	// Ping Redis to test the connection
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	// Set a key-value pair in Redis
	err = client.Set("name", "John Doe", 0).Err()
	if err != nil {
		log.Fatal("Failed to set key-value pair in Redis:", err)
	}
	fmt.Println("Key-value pair set in Redis")

	// Get the value for a key from Redis
	name, err := client.Get("name").Result()
	if err != nil {
		log.Fatal("Failed to get value from Redis:", err)
	}
	fmt.Println("Value for key 'name' in Redis:", name)

	// Increment a counter in Redis
	err = client.Incr("counter").Err()
	if err != nil {
		log.Fatal("Failed to increment counter in Redis:", err)
	}
	fmt.Println("Counter incremented in Redis")

	// Get the current value of the counter from Redis
	counter, err := client.Get("counter").Int64()
	if err != nil {
		log.Fatal("Failed to get counter value from Redis:", err)
	}
	fmt.Println("Current counter value in Redis:", counter)

	// Close the Redis connection
	err = client.Close()
	if err != nil {
		log.Fatal("Failed to close Redis connection:", err)
	}
	fmt.Println("Redis connection closed")
}
