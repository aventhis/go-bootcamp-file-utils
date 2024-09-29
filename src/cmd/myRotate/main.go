package main

import (
	"flag"
	"fmt"
	"github.com/aventhis/go-bootcamp-file-utils/src/internal/service"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	archiveDir := flag.String("a", "", "Directory to store archive files")
	flag.Parse()

	files := flag.Args()
	if len(files) < 1 {
		fmt.Println("Usage: myRotate [-a <archive_directory>] <log_file1> [<log_file2> ...]")
		os.Exit(1)
	}
	// Проверяем, существует ли выходная директория, и создаем её, если не существует
	if *archiveDir != "" {
		if err := os.MkdirAll(*archiveDir, os.ModePerm); err != nil {
			fmt.Println("Ошибка при создании директории")
			os.Exit(1)
		}
	}
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			// Проверяем, является ли файл лог-файлом (.log)
			if filepath.Ext(file) != ".log" {
				fmt.Printf("Пропускаем файл %s: допустимы только файлы .log\n", file)
				return
			}

			// Архивируем файл
			if err := service.ArchiveFile(file, *archiveDir); err != nil {
				fmt.Printf("Error archiving file %s: %s\n", file, err)
			}
		}(file)
	}
	wg.Wait()
}
