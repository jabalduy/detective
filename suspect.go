package main

import (
	"bufio"
	"fmt"
)

type Suspect struct {
	Name  string
	Age   int
	About string
	Alibi string
}

func CreateSuspects() []Suspect {
	susCollect := []Suspect{
		{
			Name:  "👷🏽‍♂️ Джозеф",
			Age:   54,
			About: "Строитель, необщительный, коренастый",
			Alibi: "Был на работе",
		},
		{
			Name:  "👩🏻‍🍳 Моника",
			Age:   36,
			About: "Повар, качает права, низкая",
			Alibi: "Была с семьей",
		},
		{
			Name:  "👨🏼‍🎓 Томас",
			Age:   19,
			About: "Студент, относится несерьезно, высокий",
			Alibi: "Был на экзамене",
		},
		{
			Name:  "👩🏻‍💼 Карен",
			Age:   49,
			About: "Бухгалтер, хочет скорее уйти, среднее телосложение",
			Alibi: "Была на выставке",
		},
		{
			Name:  "👵🏼 Лилиан",
			Age:   72,
			About: "На пенсии, за правосудие, маленькая",
			Alibi: "Была на рынке",
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
		choiceNum, ok := ReadNumber(scanner)
		if !ok {
			continue
		}

		// показать инфу
		if choiceNum > 0 && choiceNum <= len(game.Suspects) {
			fmt.Println()
			fmt.Printf("%s, %d\n", game.Suspects[choiceNum-1].Name, game.Suspects[choiceNum-1].Age)
			fmt.Println()
			fmt.Println("Описание:")
			fmt.Println(game.Suspects[choiceNum-1].About)
			fmt.Println()
			fmt.Println("Алиби:")
			fmt.Println(game.Suspects[choiceNum-1].Alibi)
			fmt.Println()
			return
		} else if choiceNum == 0 {
			return
		} else {
			fmt.Println("Такого подозреваемого нет")
			return
		}
	}
}
