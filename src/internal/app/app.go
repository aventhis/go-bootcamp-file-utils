package app

import (
	"fmt"
	"os"
)

//func parseFlag() (bool, bool, bool) {
//	flagFile := flag.Bool("f", false, "Флаг для отображения файлов")
//	flagDir := flag.Bool("d", false, "Флаг для отображения директорий")
//	flagLink := flag.Bool("sl", false, "Флаг для отображения символических ссылок")
//	flag.Parse()
//	if !*flagFile && !*flagDir && !*flagLink {
//		*flagFile = true
//		*flagDir = true
//		*flagLink = true
//	}
//	return *flagFile, *flagDir, *flagLink
//}

func MyFind() {
	//flagF, flagD, flagSL := parseFlag()
	if len(os.Args) < 2 {
		fmt.Println("необходимо указать директорию\n Usage: ~$ ./myFind /foo")
		os.Exit(1)
	}
	directoryPath := os.Args[1]
	fmt.Println(directoryPath)
}
