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

		// ИСТОРИЯ

		PrintIntro()

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

func PrintIntro() {
	fmt.Println("\n" +
		"========================================\n" +
		"              ДЕЛО №17\n" +
		"        «ПОСЛЕДНИЙ ЭКЗЕМПЛЯР»\n" +
		"========================================\n" +
		"\n" +
		"В 21:10 в библиотеке собственного дома найден\n" +
		"мёртвым Эдвард Вейл, 67 лет — состоятельный\n" +
		"коллекционер редких книг.\n" +
		"\n" +
		"Причина смерти — сильный удар по затылку.\n" +
		"Рядом с телом лежал окровавленный бронзовый\n" +
		"подсвечник.\n" +
		"\n" +
		"Из запертой витрины исчезло первое издание\n" +
		"«Чёрного сада» стоимостью около £180 000.\n" +
		"Больше ничего ценного из дома не пропало.\n" +
		"\n" +
		"----------------------------------------\n" +
		"             ИЗВЕСТНО О ДЕЛЕ\n" +
		"----------------------------------------\n" +
		"\n" +
		"Джозеф — строитель, проводивший ремонт в доме Эдварда.\n" +
		"Моника — повар Эдварда, работала в доме вечером.\n" +
		"Томас — студент, одногруппник внучки Эдварда, Эмили.\n" +
		"Карен — бухгалтер, вела финансы Эдварда.\n" +
		"Лилиан — соседка и давняя подруга семьи.\n" +
		"\n" +
		"Следов явного взлома не обнаружено.\n" +
		"Точное время смерти пока неизвестно.\n" +
		"Показания подозреваемых ещё предстоит проверить.\n" +
		"\n" +
		"Осмотрите дом. Поговорите с подозреваемыми.\n" +
		"Восстановите события вечера.\n" +
		"\n" +
		"========================================")
	fmt.Println()
}
