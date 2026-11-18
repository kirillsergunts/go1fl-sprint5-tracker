// Пакет actioninfo выводит информацию о любых активностях,
// которые умеют разбирать данные и формировать отчёт.
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — интерфейс: любой тип, у которого есть оба метода
// с такими сигнатурами, автоматически ему удовлетворяет.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info разбирает каждую строку из dataset и выводит отчёт.
// Ошибки логируются, а обработка переходит к следующей строке.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(info)
	}
}
