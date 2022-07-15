//bufio.go shows how buffer in golang works!
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a name: ")
	reader := bufio.NewReader(os.Stdin)              // СЮДА ВСТАВЛЯЕШЬ ПАРАМЕТР, КАК И В СЛУЧАЕ С strings.NewReplacer
	elementFromConsole, _ := reader.ReadString('\n') // а не read
	fmt.Println("You have entered name:", elementFromConsole)
}
