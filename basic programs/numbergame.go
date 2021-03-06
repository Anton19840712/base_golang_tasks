// Ошибки, допущенные мною при написании этой программы:
// bufferReader.ReadString('\n') а не NewReaderString('\n')
// Добавлял лишние пакеты по предположению, памяти, пытаясь сразу предугадать и написать всю логику - пиши и тестируй логику по кускам
// Не везде, но попропускал запятые в println между передаваемыми параметрами
// strconv - забыл как пишется название пакета
// при использовании if else - else писал на на одном уровне с закрывающей скобкой от if

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	guessNumber := rand.Intn(100) + 1

	fmt.Println()
	fmt.Println("Program generated a number from 1 to 100.\nTry to guess what the number is.")
	fmt.Println("You have only 3 attempts to guess, what the number is.\nSo, let's go: insert a number from 1 to 100, please...")

	for attempts := 0; attempts < 3; attempts++ {
		reader := bufio.NewReader(os.Stdin)
		stringFromConsole, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("Error occured while trying to read a string from console.")
		}

		stringFromConsoleTrimmed := strings.TrimSpace(stringFromConsole)

		numberFromString, err := strconv.Atoi(stringFromConsoleTrimmed)

		if err != nil {
			log.Fatal("Error occured while trying to parse a string from console to integer.")
		}

		fmt.Println("The number, that was inserted is: ", numberFromString)

		if numberFromString < guessNumber {
			fmt.Println("The number to guess is more, than was inserted.")
		} else if numberFromString > guessNumber {
			fmt.Println("The number to guess is less, than was inserted.")
		} else {
			fmt.Println("Congratulations, you guess the number, thar was generated by this console program.")
			break
		}
	}
	fmt.Println("The number was", guessNumber)
}
