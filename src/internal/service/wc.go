package service

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

func ParseFlagMyWc() (bool, bool, bool) {
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

func ProcessFile(filePath string, countFunc func(*os.File) (int, error), wg *sync.WaitGroup) {
	defer wg.Done()

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

func CountWords(file *os.File) (int, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return -1, errors.New("ошибка чтения файла")
	}

	words := strings.Fields(string(data))
	return len(words), nil
}

func CountStr(file *os.File) (int, error) {
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

func CountChar(file *os.File) (int, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return -1, errors.New("ошибка чтения файла")
	}
	return utf8.RuneCountInString(string(data)), nil
}
