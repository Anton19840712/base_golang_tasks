package main

import (
	"fmt"
	"strings"
)

func main() {

	initialSentence := "   Antonio, your brother is going mad...   "
	initialSentenceTrimed := strings.TrimSpace(initialSentence)
	fmt.Println(initialSentence)
	fmt.Println(initialSentenceTrimed)
}
