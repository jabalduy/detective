package main

import (
	"bufio"
	"detective/internal/game"
	"fmt"
)

func CreateSuspects() []game.Suspect {
	susCollect := []game.Suspect{
		{
			Name:  "👷🏽‍♂️ Джозеф",
			Age:   54,
			About: "Строитель. Последние несколько недель ремонтировал дом Эдварда.",
			SusID: 1,
		},
		{
			Name:  "👩🏻‍🍳 Моника",
			Age:   36,
			About: "Повар Эдварда. В вечер убийства работала в доме.",
			SusID: 2,
		},
		{
			Name:  "👨🏼‍🎓 Томас",
			Age:   19,
			About: "Студент. Одногруппник внучки Эдварда, Эмили",
			SusID: 3,
		},
		{
			Name:  "👩🏻‍💼 Карен",
			Age:   49,
			About: "Бухгалтер Эдварда. Вела его финансовые дела.",
			SusID: 4,
		},
		{
			Name:  "👵🏼 Лилиан",
			Age:   64,
			About: "Соседка и давняя подруга семьи. Живёт напротив.",
			SusID: 5,
		},
	}
	return susCollect
}

func ChooseSus(game *game.GameState, scanner *bufio.Scanner) {
	for {

		// текст выбрать подозреваемого
		fmt.Println()
		fmt.Println("=========")
		fmt.Println("Подозреваемые:")
		fmt.Println("=========")
		fmt.Println()
		for i, sus := range game.Suspects {
			fmt.Printf("%d. %s, %d\n", i+1, sus.Name, sus.Age)
		}
		fmt.Println()
		fmt.Println("0. Назад")
		fmt.Println()
		fmt.Println("Выберите подозреваемого:")

		// получить число
		choiceSus, ok := ReadNumber(scanner)
		if !ok {
			continue
		}

		// показать инфу
		if choiceSus > 0 && choiceSus <= len(game.Suspects) {

			ChooseQuestion(game, choiceSus, scanner)

		} else if choiceSus == 0 {
			ClearScreen()
			return
		} else {
			fmt.Println("Такого подозреваемого нет")
			ClearScreen()
			continue
		}
	}
}
