package main

import (
	"bufio"
	"fmt"
)

type Suspect struct {
	Name   string
	Age    int
	About  string
	Alibi  string
	SusID  int
	IsClue bool
}

func CreateSuspects() []Suspect {
	susCollect := []Suspect{
		{
			Name:  "👷🏽‍♂️ Джозеф",
			Age:   54,
			About: "Строитель. Последние несколько недель ремонтировал дом Эдварда.",
			Alibi: "Был на работе",
			SusID: 1,
		},
		{
			Name:  "👩🏻‍🍳 Моника",
			Age:   36,
			About: "Повар Эдварда. В вечер убийства работала в доме.",
			Alibi: "Была с семьей",
			SusID: 2,
		},
		{
			Name:  "👨🏼‍🎓 Томас",
			Age:   19,
			About: "Студент. Одногруппник внучки Эдварда, Эмили",
			Alibi: "Был на экзамене",
			SusID: 3,
		},
		{
			Name:  "👩🏻‍💼 Карен",
			Age:   49,
			About: "Бухгалтер Эдварда. Вела его финансовые дела.",
			Alibi: "Была на выставке",
			SusID: 4,
		},
		{
			Name:  "👵🏼 Лилиан",
			Age:   64,
			About: "Соседка и давняя подруга семьи. Живёт напротив.",
			Alibi: "Была на рынке",
			SusID: 5,
		},
	}
	return susCollect
}

func ChooseSus(game *GameState, scanner *bufio.Scanner) {
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
			return
		} else {
			fmt.Println("Такого подозреваемого нет")
			continue
		}
	}
}
