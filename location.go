package main

import (
	"bufio"
	"fmt"
)

type Location struct {
	Name  string
	About string
}

func CreateLocations() []Location {
	locCollect := []Location{
		{
			Name:  "Гостиная",
			About: "Пусто",
		},
		{
			Name:  "Кухня",
			About: "Пусто",
		},
		{
			Name:  "Библиотека",
			About: "Пусто",
		},
	}
	return locCollect
}

func ChooseLocations(locCol []Location, scanner *bufio.Scanner) {
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
		choiceNum, ok := ReadNumber(scanner)
		if !ok {
			continue
		}

		// показать инфу
		if choiceNum > 0 && choiceNum <= len(locCol) {
			fmt.Println()
			fmt.Println(locCol[choiceNum-1].Name)
			fmt.Println(locCol[choiceNum-1].About)
			fmt.Println()
			return
		} else if choiceNum == 0 {
			return
		} else {
			fmt.Println("Такого места нет.")
			return
		}
	}
}
