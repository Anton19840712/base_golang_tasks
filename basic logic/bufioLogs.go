//bufioLogs.go: here we use if nil expression (means if there is no mistakes by default) and log package

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("Enter any name, please: ")
	reader := bufio.NewReader(os.Stdin)
	bufioReadResult, err := reader.ReadString('\n')

	if err != nil {//лучше писать таким образом: != nil, а не == nil
		log.Fatal(err) // просто err - не достаточно.
		//fmt.Println(err)// this is useless, result with or without this is the same.
	}
	
	fmt.Println("You have entered name: ", bufioReadResult)	
}
