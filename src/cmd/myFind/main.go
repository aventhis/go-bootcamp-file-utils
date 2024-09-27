// Поиск всех файлов/каталогов/символических ссылок рекурсивно в каталоге
package main

import (
	"flag"
	"fmt"
	"github.com/aventhis/go-bootcamp-file-utils/src/internal/service"
	"os"
)

func main() {
	flagF, flagD, flagSL, flagEXT := service.ParseFlagFind()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Необходимо указать ровно одну директорию\nUsage: ~$ ./myFind /foo")
		fmt.Println("или если флаг ext, то указать расширение файла\nUsage: ~$ ./myFind -f -ext 'go' /go")
		os.Exit(1)
	}

	root := args[0]

	err := service.WalkDirectory(root, flagF, flagD, flagSL, flagEXT)
	if err != nil {
		fmt.Printf("Ошибка обхода директории: %v\n", err)
	}

}
