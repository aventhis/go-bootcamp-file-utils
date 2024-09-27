package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flagA := flag.String("a", "", "flag indicating the directory for archives")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Usage: myRotate [-a <archive_directory>] <log_file1> [<log_file2> ...]")
		os.Exit(1)
	}

	fmt.Println(*flagA)
}
