package main

import (
	"flag"
	"fmt"
	"github.com/aventhis/go-bootcamp-file-utils/src/internal/service"
	"os"
	"sync"
)

func main() {
	flagW, flagL, flagM := service.ParseFlagMyWc()

	files := flag.Args()

	if len(files) == 0 {
		fmt.Println("Usage: ./myWc -m input4.txt input5.txt input6.txt")
		os.Exit(1)
	}

	//переопределяем функцию
	var countFunc func(file *os.File) (int, error)
	if flagW {
		countFunc = service.CountWords
	} else if flagL {
		countFunc = service.CountStr
	} else if flagM {
		countFunc = service.CountChar
	} else {
		countFunc = service.CountWords
	}

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go service.ProcessFile(file, countFunc, &wg)
	}
	wg.Wait()
}
