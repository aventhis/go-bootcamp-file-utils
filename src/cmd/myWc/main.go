package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	flagW := flag.Bool("w", false, "Флаг для подсчета слов")
	flagL := flag.Bool("l", false, "Флаг для подсчета строк")
	flagM := flag.Bool("m", false, "Флаг для подсчета символов")

	flag.Parse()

	if *flagW && *flagL || *flagW && *flagM || *flagL && *flagM {
		fmt.Println("Ошибка: может бытьь указан только один флаг")
		fmt.Println("Usage: ./myWc -m input4.txt")
		os.Exit(1)
	}

	files := flag.Args()

	if len(files) == 0 {
		fmt.Println("Usage: ./myWc -m input4.txt input5.txt input6.txt")
		os.Exit(1)
	}

	fmt.Println(files)

	//переопределяем функцию
	var countFunc func(file *os.File) (int, error)
	if *flagW {
		countFunc = countWords
	} else if *flagL {
		countFunc = countStr
	} else if *flagM {
		countFunc = countChar
	}

	for _, file := range files {
		processFile(file, countFunc)
	}

}

func processFile(filePath string, countFunc func(*os.File) (int, error)) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Ошибка закрытия файла", err)
			return
		}
	}()

	count, err := countFunc(file)
	if err != nil {
		fmt.Println("Ошибка обработки файла", filePath, err)
		return
	}

	fmt.Printf("%d\t%s\n", count, filePath)

}

func countWords(file *os.File) (int, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return -1, errors.New("ошибка чтения файла")
	}

	words := strings.Fields(string(data))
	return len(words), nil
}

func countStr(file *os.File) (int, error) {
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return count, nil
}

func countChar(file *os.File) (int, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return -1, errors.New("ошибка чтения файла")
	}
	return utf8.RuneCountInString(string(data)), nil
}
