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

//// Рекурсивный обход директорий с использованием filepath.Walk
//err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
//	if err != nil {
//		fmt.Println("Ошибка при чтении файла:", err)
//		return nil
//	}
//
//	// Проверка типа объекта и вывод в зависимости от флагов
//	if info.Mode()&os.ModeSymlink != 0 {
//		// Это символическая ссылка
//		if flagSL {
//			linkTarget, err := os.Readlink(path)
//			if err != nil {
//				fmt.Printf("%s -> [broken]\n", path)
//			} else {
//				fmt.Printf("%s -> %s\n", path, linkTarget)
//			}
//		}
//	} else if info.IsDir() {
//		// Это директория
//		if flagD {
//			fmt.Println(path)
//		}
//	} else {
//		// Это обычный файл
//		if flagF {
//			fmt.Println(path)
//		}
//	}
//	return nil
//})

//if err != nil {
//	fmt.Println("Ошибка при обходе директорий:", err)
//}
