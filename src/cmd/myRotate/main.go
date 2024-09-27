package main

import (
	"flag"
	"fmt"
	"github.com/aventhis/go-bootcamp-file-utils/src/internal/service"
	"os"
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

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			if err := service.ArchiveFile(file, *archiveDir); err != nil {
				fmt.Printf("Error archiving file %s: %s\n", file, err)
			}
		}(file)
	}
	wg.Wait()
}
