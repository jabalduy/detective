package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("ДЕЛО №17")
		fmt.Println("")
		fmt.Println("1. Исследовать место")
		fmt.Println("2. Допросить подозреваемого")
		fmt.Println("3. Посмотреть улики")
		fmt.Println("4. Обвинить")
		fmt.Println("5. Выйти")
		fmt.Println("")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода!")
			continue
		}

		text := scanner.Text()
		if len(text) == 0 {
			fmt.Println("Пустой ввод!")
			continue
		}

		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Ошибка преобразования!", err)
			continue
		}

		switch num {
		case 1:
			fmt.Println()
			fmt.Println("Выбрано 1 - Исследовать место")
			fmt.Println()
		case 2:
			fmt.Println()
			fmt.Println("Выбрано 2 - Допросить подозреваемого")
			fmt.Println()
		case 3:
			fmt.Println()
			fmt.Println("Выбрано 3 - Посмотреть улики")
			fmt.Println()
		case 4:
			fmt.Println()
			fmt.Println("Выбрано 4 - Обвинить")
			fmt.Println()
		case 5:
			fmt.Println()
			fmt.Println("Выбрано 5 - Выйти")
			fmt.Println()
			return
		default:
			fmt.Println()
			fmt.Println("Неизвестная команда")
			fmt.Println()
		}
	}
}
