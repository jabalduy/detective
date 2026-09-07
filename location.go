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

func ChooseLocations(locCol []Location, cluesCol []Clue, objCol []Object, scanner *bufio.Scanner) {

	for {

		// текст выбрать  локации
		fmt.Println()
		fmt.Println("Локации:")
		for i, loc := range locCol {
			fmt.Printf("%d. %s\n", i+1, loc.Name)
		}
		fmt.Println()
		fmt.Println("0. Назад")
		fmt.Println()
		fmt.Println("Выберите локацию:")

		// получить число
		choiceLoc, ok := ReadNumber(scanner)
		if !ok {
			continue
		}

		// показать инфу
		if choiceLoc > 0 && choiceLoc <= len(locCol) {
			fmt.Println()
			fmt.Printf("%s:\n", locCol[choiceLoc-1].Name)
			fmt.Println(locCol[choiceLoc-1].About)

			ChooseObjects(objCol, cluesCol, scanner, choiceLoc)

			fmt.Println()
			return
		} else if choiceLoc == 0 {
			return
		} else {
			fmt.Println("Такого места нет.")
			return
		}
	}
}
