package app

import (
	"flag"
	"fmt"
	"github.com/aventhis/go-bootcamp-file-utils/src/internal/service"
	"os"
)

func MyFind() {
	flagF, flagD, flagSL := service.ParseFlag()
	if flagF {
		fmt.Println("file")
	}
	if flagD {
		fmt.Println("dir")
	}
	if flagSL {
		fmt.Println("sl")
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Необходимо указать директорию\nUsage: ~$ ./myFind /foo")
		os.Exit(1)
	}

	directoryPath := args[0]
	fmt.Println(directoryPath)
}
