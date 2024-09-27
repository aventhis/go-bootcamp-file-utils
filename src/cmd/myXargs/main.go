package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: myXargs <command> [args...]")
		os.Exit(1)
	}

	// Читаем команду и аргументы из командной строки
	command := os.Args[1]
	args := os.Args[2:]

	// Читаем данные из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		args = append(args, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения из Stdin:", err)
		os.Exit(1)
	}

	// Выполняем команду с переданными аргументами
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Ошибка выполнения команды:", err)
		os.Exit(1)
	}
}
