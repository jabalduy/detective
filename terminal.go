package main

import (
	"bufio"
	"fmt"
)

const (
	Green = "\033[32m"
	Red   = "\033[31m"
	Gold  = "\033[33m"
	Reset = "\033[0m"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Pause(scanner *bufio.Scanner) {
	fmt.Println()
	fmt.Println("Нажмите Enter, чтобы продолжить...")
	scanner.Scan()
}
