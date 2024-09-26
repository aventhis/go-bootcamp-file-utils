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
		fmt.Println("Ошибка: флаг -ext можно использовать только с флагом -f.")
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

func WalkDirectory(root string, flagF, flagD, flagSL bool, flagEXT string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Пропускаем файлы или директории, к которым нет доступа
			fmt.Printf("Ошибка при доступе к %s: %v\n", path, err)
			return nil // Продолжаем обход
		}

		if flagD && info.IsDir() {
			fmt.Println(path)
		}

		// Обработка файлов
		if flagF && info.Mode().IsRegular() {
			if flagEXT != "" {
				// Проверяем расширение файла
				if filepath.Ext(path) == "."+flagEXT {
					fmt.Println(path)
				}
			} else {
				// Если расширение не указано, выводим все файлы
				fmt.Println(path)
			}
		}

		// Обработка символических ссылок
		if flagSL && (info.Mode()&os.ModeSymlink != 0) {
			link, err := os.Readlink(path)
			if err != nil {
				fmt.Printf("%s -> [broken]\n", path)
			} else {
				if _, err := os.Stat(link); os.IsNotExist(err) {
					fmt.Printf("%s -> [broken]\n", path)
				} else {
					fmt.Printf("%s -> %s\n", path, link)
				}
			}
		}

		return nil // Продолжаем обход
	})
}
