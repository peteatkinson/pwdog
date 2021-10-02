package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// USAGE: go run main.go -s 70
	size := flag.Int("s", 15, "Size of password")
	flag.Parse()
	s := *size

	arr := shuffleChars()

	if s >= len(arr) {
		fmt.Printf("Please keep password no larger than %d", len(arr))
	} else {
		// create PWD chunk from min=0 to max=s
		password := string(arr[0:s])
		fmt.Println(password)
	}
}


func shuffleChars() []byte {
	// shuffle possible a-Za-z|SPECIAL characters with some unique seed
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*-"
	arr := []byte(characters)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(left, right int) { 
		arr[left], 
		arr[right] = arr[right],
		arr[left] 
	})
	return arr
}