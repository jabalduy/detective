package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func ReadNumber(scanner *bufio.Scanner) (int, bool) {
	valid := true

	if ok := scanner.Scan(); !ok {
		fmt.Println("Ошибка ввода!")
		ClearScreen()
		valid = false
	}

	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("Пустой ввод!")
		ClearScreen()
		valid = false
	}

	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Ошибка преобразования!")
		ClearScreen()
		valid = false
	}

	return num, valid
}
