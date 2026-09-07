package main

import (
	"bufio"
	"fmt"
)

type Clue struct {
	Name  string
	About string
	Found bool
}

func CreateClue() []Clue {
	clueCollect := []Clue{
		{
			Name:  "Мокрый след",
			About: "",
			Found: false,
		},
		{
			Name:  "Разбитые часы",
			About: "",
			Found: false,
		},
		{
			Name:  "Записка",
			About: "",
			Found: false,
		},
		{
			Name:  "Чужой ключ",
			Found: false,
		},
		{
			Name:  "Пятно краски",
			About: "Найдено в библиотеке",
			Found: true,
		},
		{
			Name:  "Чек",
			About: "Найден у Джона",
			Found: true,
		},
	}
	return clueCollect
}

func ShowClues(cluesCol []Clue, scanner *bufio.Scanner) {
	for {

		// показать найденные улики
		fmt.Println()
		fmt.Println("Улики:")
		count := 0
		for _, clue := range cluesCol {
			if clue.Found {
				count += 1
				fmt.Printf("%d. %s\n", count, clue.Name)
				fmt.Println(clue.About)
				fmt.Println()
			}
		}
		if count == 0 {
			fmt.Println("Улики пока не найдены.")
		}

		fmt.Println("0. Назад")

		// получить число для выхода
		exitNum, ok := ReadNumber(scanner)
		if !ok {
			continue
		}

		// выход
		if exitNum == 0 {
			return
		} else {
			continue
		}
	}
}
