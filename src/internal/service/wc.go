package service

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

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

	return count, scanner.Err()
}

func CountChar(file *os.File) (int, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return -1, errors.New("ошибка чтения файла")
	}
	return utf8.RuneCountInString(string(data)), nil
}
