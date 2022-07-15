package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	currentUnixValue := time.Now().Unix()
	//random := rand.Seed(currentUnixValue)
	rand.Seed(currentUnixValue)
	randomValue := rand.Intn(100)
	fmt.Println(randomValue)
}
