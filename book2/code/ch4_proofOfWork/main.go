package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"time"
)

func main() {
	pow("Tanmay Bakshi + Baheer Kamal", 24)
}

func pow(prefix string, bitLength int) {
	start := time.Now()

	totalHashesProcessed := 0
	seed := uint64(time.Now().Local().UnixNano())
	randomBytes := make([]byte, 20)
	randomBytes = append([]byte(prefix), randomBytes...)
	for {
		totalHashesProcessed++
		seed = RandomString(randomBytes, len(prefix), seed)
		if Hash(randomBytes, bitLength) {
			fmt.Println(string(randomBytes))
			break
		}
		if totalHashesProcessed%10000000 == 0 {
			fmt.Println("hashes processed:", humanize.Comma(int64(totalHashesProcessed)))
		}
	}
	end := time.Now()

	fmt.Println("time:", end.Sub(start).Seconds())
	fmt.Println("process", humanize.Comma(int64(totalHashesProcessed)), "hashes")
	fmt.Printf("hash rate: %s hashes per second\n", humanize.Comma(int64(float64(totalHashesProcessed)/end.Sub(start).Seconds())))
}

func Hash(data []byte, bitLength int) bool {
	for i := 0; i < bitLength; i++ {
		if data[i/8]&(1<<uint(i%8)) == 0 {
			return false
		}
	}
	return true
}

var charactes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomNumber(seed uint64) uint64 {
	seed ^= seed << 21
	seed ^= seed >> 35
	seed ^= seed << 4
	return seed
}

func RandomString(data []byte, length int, seed uint64) uint64 {
	for i := 0; i < length; i++ {
		seed = RandomNumber(seed)
		data[i] = charactes[seed%uint64(len(charactes))]
	}
	return seed
}
