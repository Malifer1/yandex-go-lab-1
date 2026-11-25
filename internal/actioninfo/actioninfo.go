package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает структуры, умеющие разбирать данные и формировать строку с информацией.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info разбирает слайс строк и выводит информацию по каждой активности.
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
