package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	susCollect := CreateSuspects()
	clueCollect := CreateClue()
	locCollect := CreateLocations()
	objCollect := CreateObjects()

	for {
		// Текст меню
		fmt.Println("")
		fmt.Println("ДЕЛО №17")
		fmt.Println("")
		fmt.Println("1. Исследовать комнаты")
		fmt.Println("2. Допросить подозреваемого")
		fmt.Println("3. Посмотреть улики")
		fmt.Println("4. Обвинить")
		fmt.Println("5. Выйти")
		fmt.Println("")

		num, ok := ReadNumber(scanner)
		if !ok {
			continue
		}

		switch num {
		case 1:

			ChooseLocations(locCollect, clueCollect, objCollect, scanner)

		case 2:

			ChooseSus(susCollect, scanner)

		case 3:

			ShowClues(clueCollect, scanner)

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
