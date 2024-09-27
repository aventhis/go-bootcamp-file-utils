package service

import (
	"flag"
	"fmt"
	"os"
)

// ParseFlagFind - парсинг флагов командной строки для myFind
func ParseFlagFind() (bool, bool, bool, string) {
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

// ParseFlagWc - парсинг флагов командной строки для myWc
func ParseFlagWc() (bool, bool, bool) {
	flagW := flag.Bool("w", false, "Флаг для подсчета слов")
	flagL := flag.Bool("l", false, "Флаг для подсчета строк")
	flagM := flag.Bool("m", false, "Флаг для подсчета символов")

	flag.Parse()

	if *flagW && *flagL || *flagW && *flagM || *flagL && *flagM {
		fmt.Println("Ошибка: может быть указан только один флаг")
		fmt.Println("Usage: ./myWc -m input4.txt")
		os.Exit(1)
	}
	return *flagW, *flagL, *flagM
}
