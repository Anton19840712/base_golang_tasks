// ошибки, которые я допустил при написании этой программы:
// написал os.Stdin c маленькой буквы - а надо с большой, потому что это экспортируемый тип, но я исходя из чего я был должен понять, что он экспортируемый? Потому что из внешнего пакета.
// написал Trim вместо TrimSpace
// написал ParseToFloat64 вместо просто ParseFloat
// забыл, что что ParseFloat, так же, как и ReadString возвращает err
// итак:
// ReadString
// ParseFloat
// в ParseFloat забыл указать разрядность

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a score, and this programm will find out a passing status of your exam: ")
	bufferReader := bufio.NewReader(os.Stdin) // нужно Stdin с большой буквы писать
	element, err := bufferReader.ReadString('\n')

	var grade float64
	statusPassed := "passed"
	statusFailed := "failed"

	if err != nil {
		log.Fatal(err)
	} else { //здесь не нужно else - потому что если программа не завершается - то она сама по себе продолжит свое исполнение
		aspect := strings.TrimSpace(element)        //Trim не возвращает вторую переменную, но требует вторую на вход? Нужно и можно использовать здесь TrimSpace. Trim тоже есть в golang, но он работает как Replace in C#. Указываешь строку и символ.
		grade, err = strconv.ParseFloat(aspect, 64) //не забывай указывать точность или примерный размер выделения памяти и не ParseToFloat64, а функция просто ParseFloat
		//их всего 4: ParseBool, ParseFloat, ParseInt, and ParseUint
	}

	if err != nil {
		log.Fatal(err)
	}

	if grade > 60 {
		fmt.Println("Congratulations, the status of your exam is: ", statusPassed)
	} else {
		fmt.Println("Status of your exam is: ", statusFailed)
	}
}
