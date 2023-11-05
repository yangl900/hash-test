package main

import (
	"fmt"
	"time"

	"github.com/minio/sha256-simd"
)

const (
	GB   = 1024 * 1024 * 1024
	SIZE = 2 * GB
)

func main() {
	buffer := make([]byte, SIZE)
	for i := 0; i < SIZE; i++ {
		buffer[i] = 0
	}

	for i := 0; i < 5; i++ {
		// Create a new SHA256 hash.
		hash := sha256.New()
		start := time.Now()

		// Write the string to the hash.
		hash.Write(buffer)

		// Get the hash value.
		hashValue := hash.Sum(nil)

		elapsed := time.Since(start)
		fmt.Printf("Hashing took %s\n", elapsed)
		// Print the hash value.
		fmt.Println(hashValue)
	}
}
