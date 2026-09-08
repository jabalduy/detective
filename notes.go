package main

import (
	"bufio"
	"fmt"
)

func ShowCaseFile(game *GameState, scanner *bufio.Scanner) {
	for {
		fmt.Println()
		fmt.Println("=======================")
		fmt.Println("	ДОСЬЕ ДЕЛА")
		fmt.Println("=======================")
		fmt.Println()

		// ПРОГРЕСС

		objStat := 0
		for _, obj := range game.Objects {
			if obj.Searched {
				objStat += 1
			}
		}

		clueStat := 0
		for _, clue := range game.Clues {
			if clue.Found {
				clueStat += 1
			}
		}

		dialStat := 0
		for _, dial := range game.Dialogues {
			if dial.Asked {
				dialStat += 1
			}
		}

		fmt.Println("ПРОГРЕСС")
		fmt.Println()
		fmt.Println("Исследовано предметов:", objStat, "/", len(game.Objects))
		fmt.Println("Найдено улик:", clueStat, "/", len(game.Clues))
		fmt.Println("Задано вопросов:", dialStat, "/", len(game.Dialogues))

		// УЛИКИ

		ShowClues(game, scanner)

		// ПОКАЗАНИЯ
		fmt.Println("ПОКАЗАНИЯ")
		fmt.Println()
		for i := 0; i < len(game.Suspects); i++ {
			hasStatements := false

			for _, dialogue := range game.Dialogues {
				if dialogue.Asked && dialogue.SusID == i+1 {

					if !hasStatements {
						fmt.Printf("%s:\n", game.Suspects[i].Name)
						hasStatements = true
					}

					fmt.Printf("- %s\n", dialogue.Fact)
				}
			}

			if hasStatements {
				fmt.Println()
			}
		}

		// ВЫХОД

		fmt.Println()
		fmt.Println("0. Назад")
		fmt.Println()
		choiceNum, ok := ReadNumber(scanner)
		if !ok {
			fmt.Println("Неверный ввод")
			continue
		}

		if choiceNum == 0 {
			return
		}

	}
}
