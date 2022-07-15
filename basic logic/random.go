package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.Intn(100) // принимает только одно значение
	fmt.Println(random)
}
