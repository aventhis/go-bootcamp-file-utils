package service

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Finder(directoryPath string, flagF, flagD, flagSL bool) {
	dir, err := os.Open(directoryPath)
	if err != nil {
		fmt.Println("Не удалось открыть директорию", directoryPath)
		os.Exit(1)
	}

	defer func() {
		if err := dir.Close(); err != nil {
			fmt.Println("Ошибка при закрытии директории:", err)
		}
	}()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Ошибка чтения содержимого директории", err)
	}

	for _, file := range files {
		fullPath := filepath.Join(directoryPath, file.Name())

		fileInfo, err := os.Lstat(fullPath)
		if err != nil {
			fmt.Println("Ошибка получения информации о файле:", err)
			continue
		}

		if fileInfo.Mode()&os.ModeSymlink != 0 {
			if flagSL {
				// Чтение цели символической ссылки
				linkTarget, err := os.Readlink(fullPath)
				if err != nil {
					fmt.Printf("%s -> [broken]\n", fullPath)
				} else {
					// Проверка, существует ли целевой объект ссылки
					if _, err := os.Stat(fullPath); os.IsNotExist(err) {
						fmt.Printf("%s -> [broken]\n", fullPath)
					} else {
						fmt.Printf("%s -> %s\n", fullPath, linkTarget)
					}
				}
			}
			continue
		}

		if flagF && !fileInfo.IsDir() {
			fmt.Println(fullPath)
		}

		if flagD && fileInfo.IsDir() {
			fmt.Println(fullPath)
		}

		if fileInfo.IsDir() {
			Finder(fullPath, flagF, flagD, flagSL)
		}

	}

}

func ParseFlag() (bool, bool, bool) {
	flagFile := flag.Bool("f", false, "Флаг для отображения файлов")
	flagDir := flag.Bool("d", false, "Флаг для отображения директорий")
	flagLink := flag.Bool("sl", false, "Флаг для отображения символических ссылок")
	flag.Parse()
	if !*flagFile && !*flagDir && !*flagLink {
		*flagFile = true
		*flagDir = true
		*flagLink = true
	}
	return *flagFile, *flagDir, *flagLink
}
