//допущенные ошибки: time.Now() написал now с маленькой буквы. Понимай, что now - это часть внешнего пакета, и как правило после названия пакета пишется не объект, а метод. И так как он внешний, то он пишется с большой буквы

package main

import (
	"fmt"
	"time"
)

func main() {
	valueElement := time.Now().Unix()
	fmt.Println(valueElement)
}
