package main

import (
	"bufio"
	"fmt"
)

type Location struct {
	Name   string
	About  string
	Object Object
	LocID  int
}

func CreateLocations() []Location {
	locCollect := []Location{
		{
			Name:  "Гостиная",
			About: "Небольшая гостиная рядом с главным коридором.\nНа диване лежат учебники,\nна журнальном столике стоит недопитый чай.",
			LocID: 1,
		},
		{
			Name:  "Кухня",
			About: "На кухне осталась посуда после приготовления ужина.\nНа стене тикают старые механические часы.\nОдин из нижних шкафов закрывается неплотно.",
			LocID: 2,
		},
		{
			Name:  "Библиотека",
			About: "Большая комната с книжными шкафами до потолка.\nВозле камина лежит тело хозяина дома — Эдварда Вейла.\nНа первый взгляд следов взлома нет.",
			LocID: 3,
		},
	}
	return locCollect
}

func ChooseLocations(game *GameState, scanner *bufio.Scanner) {

	for {
		ClearScreen()

		// текст выбрать  локации
		fmt.Println()
		fmt.Println("Комнаты:")
		for i, loc := range game.Locations {
			fmt.Printf("%d. %s\n", i+1, loc.Name)
		}
		fmt.Println()
		fmt.Println("0. Назад")
		fmt.Println()
		fmt.Println("Выберите комнату:")

		// получить число
		choiceLoc, ok := ReadNumber(scanner)
		if !ok {
			fmt.Println("Неверный ввод")
			continue
		}

		// показать инфу
		if choiceLoc > 0 && choiceLoc <= len(game.Locations) {
			fmt.Println()
			fmt.Println("=======================")
			fmt.Printf("%s:\n", game.Locations[choiceLoc-1].Name)
			fmt.Println("=======================")
			fmt.Println()
			fmt.Println(game.Locations[choiceLoc-1].About)

			ChooseObjects(game, scanner, choiceLoc)

			fmt.Println()
		} else if choiceLoc == 0 {
			ClearScreen()
			return
		} else {
			ClearScreen()
			fmt.Println("Такого места нет.")
			continue
		}
	}
}
