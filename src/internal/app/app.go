package app

import (
	"flag"
	"fmt"
	"github.com/aventhis/go-bootcamp-file-utils/src/internal/service"
	"os"
)

func MyFind() {
	flagF, flagD, flagSL := service.ParseFlag()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Необходимо указать ровно одну директорию\nUsage: ~$ ./myFind /foo")
		os.Exit(1)
	}

	directoryPath := args[0]
	service.Finder(directoryPath, flagF, flagD, flagSL)

}
