package service

import (
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
		if file.IsDir() {
			Finder(fullPath, flagF, flagD, flagSL)
		}
		fmt.Println(fullPath)
	}
}
