package service

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// ParseFlag - парсинг флагов командной строки
func ParseFlag() (bool, bool, bool, string) {
	flagFile := flag.Bool("f", false, "Флаг для отображения файлов")
	flagDir := flag.Bool("d", false, "Флаг для отображения директорий")
	flagLink := flag.Bool("sl", false, "Флаг для отображения символических ссылок")

	flagEXT := flag.String("ext", "", "Флаг для отображения файлов с определенным расширением")

	flag.Parse()

	if !*flagFile && *flagEXT != "" {
		*flagEXT = ""
		fmt.Println("Usage: ~$ ./myFind -f -ext 'go' /go")
		os.Exit(1)
	}

	if !*flagFile && !*flagDir && !*flagLink {
		*flagFile = true
		*flagDir = true
		*flagLink = true
	}

	return *flagFile, *flagDir, *flagLink, *flagEXT
}

func Finder(directoryPath string, flagF, flagD, flagSL bool, flagEXT string) {
	files := readDir(directoryPath)

	if files == nil {
		return
	}

	for _, file := range files {
		fullPath := filepath.Join(directoryPath, file.Name())

		// Получение информации о файле или символической ссылке
		fileInfo, err := os.Lstat(fullPath)
		if err != nil {
			if os.IsPermission(err) {
				// Пропускаем файлы с ошибками доступа
				fmt.Println("Пропущен файл из-за недостаточных прав:", fullPath)
				continue
			}
			fmt.Println("Ошибка получения информации о файле:", err)
			continue
		}

		// Обработка символической ссылки
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			if flagSL {
				processSymlink(fullPath)
			}
			continue
		}

		// Обработка файлов и директорий
		if fileInfo.IsDir() {
			if flagD {
				fmt.Println(fullPath)
			}
			Finder(fullPath, flagF, flagD, flagSL, flagEXT)
		} else if flagF {
			if flagEXT != "" {
				ext := filepath.Ext(file.Name())
				if ext == "."+flagEXT {
					fmt.Println(fullPath)
				}
			} else {
				fmt.Println(fullPath)
			}
		}
	}

}

// readDir - открытие и чтение директории
func readDir(directoryPath string) []os.FileInfo {
	dir, err := os.Open(directoryPath)
	if err != nil {
		if os.IsPermission(err) {
			// Пропускаем директорию, если нет доступа
			fmt.Println("Пропущена директория из-за недостаточных прав:", directoryPath)
			return nil
		}
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
		if os.IsPermission(err) {
			// Пропускаем чтение содержимого директории, если нет доступа
			fmt.Println("Пропущено чтение содержимого директории из-за недостаточных прав:", directoryPath)
			return nil
		}
		fmt.Println("Ошибка чтения содержимого директории", err)
	}

	return files
}

// processSymlink - Функция для обработки символических ссылок
func processSymlink(fullPath string) {
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
