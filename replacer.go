//replacer.go replaces from one string to another one
package main

import (
	"fmt"
	"strings"
)

func main() {
	stringElement := "This is an Ironman"
	replacer := strings.NewReplacer("Ironman", "Antonio")
	changedElement := replacer.Replace(stringElement)
	fmt.Println(changedElement)
}
