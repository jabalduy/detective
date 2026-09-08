package main

import (
	"bufio"
	"fmt"
)

func ChooseAccused(game *GameState, scanner *bufio.Scanner) bool {
	for {
		fmt.Println()
		fmt.Println("Выберите обвиняемого:")
		fmt.Println()
		for i, sus := range game.Suspects {
			fmt.Printf("%d. %s\n", i+1, sus.Name)
		}

		fmt.Println()
		fmt.Println("0. Назад")
		fmt.Println()

		choiceAc, ok := ReadNumber(scanner)
		if !ok {
			ClearScreen()
			fmt.Println("Неверный ввод")
			continue
		}

		if choiceAc < 0 || choiceAc > len(game.Suspects) {
			ClearScreen()
			fmt.Println("Неверный ввод")
			continue
		}

		if choiceAc == 0 {
			ClearScreen()
			return false
		}

		ClearScreen()
		if Sure(game, choiceAc, scanner) {
			return true
		} else {
			ClearScreen()
			continue
		}
	}
}

func Sure(game *GameState, choiceAc int, scanner *bufio.Scanner) bool {
	for {
		fmt.Println()
		fmt.Printf("Вы уверены, что обвиняете %s?\n", game.Suspects[choiceAc-1].Name)
		fmt.Println()
		fmt.Println("1. ДА")
		fmt.Println("0. НЕТ")
		fmt.Println()

		choiceSure, ok := ReadNumber(scanner)
		if !ok {
			ClearScreen()
			fmt.Println("Неверный ввод")
			continue
		}

		if choiceSure < 0 || choiceSure > 1 {
			ClearScreen()
			fmt.Println("Неверный ввод")
			continue
		}

		ClearScreen()
		if choiceSure == 1 {
			ClearScreen()
			Accuse(choiceAc, game, scanner)
			return true
		}

		if choiceSure == 0 {
			ClearScreen()
			return false
		}
	}
}

func Accuse(choiceAc int, game *GameState, scanner *bufio.Scanner) {
	if KeyClues(game)+KeyDial(game)+KeyObj(game) == 7 && choiceAc == 4 {
		WinEnding()
		Pause(scanner)
	} else if KeyClues(game)+KeyDial(game)+KeyObj(game) < 7 && choiceAc == 4 {
		UnsolvedEnding()
		Pause(scanner)
	} else if KeyClues(game)+KeyDial(game)+KeyObj(game) < 7 && choiceAc != 4 {
		FailedEnding()
		Pause(scanner)
	}
}

func WinEnding() {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("              ОБВИНЕНИЕ")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Вы называете имя:")
	fmt.Println()
	fmt.Println("Карен Моррис.")
	fmt.Println()
	fmt.Println("Следствие складывается в единую картину.")
	fmt.Println()
	fmt.Println("Эдвард погиб около 20:37.")
	fmt.Println("В камине найден документ MORRIS CONSULTING,")
	fmt.Println("связанный с финансовыми махинациями.")
	fmt.Println()
	fmt.Println("Карен утверждала, что приехала позже")
	fmt.Println("и вообще не заходила в библиотеку.")
	fmt.Println()
	fmt.Println("Но её видели возле кухни раньше, чем она призналась,")
	fmt.Println("а пропавшая книга была спрятана именно там.")
	fmt.Println()
	fmt.Println("Кража была инсценировкой.")
	fmt.Println("Карен попыталась скрыть настоящий мотив убийства.")
	fmt.Println()
	fmt.Println("Несколько секунд она молчит.")
	fmt.Println()
	fmt.Println("Затем тихо говорит:")
	fmt.Println()
	fmt.Println("«Он обнаружил всё.")
	fmt.Println("На следующее утро собирался обратиться к адвокату.")
	fmt.Println("Я хотела забрать документы.")
	fmt.Println("Мы поссорились...")
	fmt.Println("Я не собиралась его убивать.»")
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println(Green + "              ДЕЛО ЗАКРЫТО" + Reset)
	fmt.Println("========================================")
	fmt.Println()
}

func UnsolvedEnding() {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("              ОБВИНЕНИЕ")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Вы обвиняете Карен Моррис.")
	fmt.Println()
	fmt.Println("Версия выглядит правдоподобно,")
	fmt.Println("но в расследовании остаются пробелы.")
	fmt.Println()
	fmt.Println("Часть важных улик не найдена,")
	fmt.Println("некоторые показания не проверены,")
	fmt.Println("а цепочка событий не доказана полностью.")
	fmt.Println()
	fmt.Println("Карен спокойно отрицает обвинение.")
	fmt.Println()
	fmt.Println("Полиции недостаточно оснований,")
	fmt.Println("чтобы предъявить ей обвинение.")
	fmt.Println()
	fmt.Println("Карен покидает дом.")
	fmt.Println()
	fmt.Println("Вы знаете, что могли быть правы.")
	fmt.Println("Но знать — недостаточно.")
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println(Red + "             НЕДОСТАТОЧНО УЛИК" + Reset)
	fmt.Println("========================================")
	fmt.Println()
}

func FailedEnding() {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("              ОБВИНЕНИЕ")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Вы делаете окончательный выбор.")
	fmt.Println()
	fmt.Println("Но собранные факты не складываются")
	fmt.Println("в доказательство вины этого человека.")
	fmt.Println()
	fmt.Println("Расследование заканчивается ошибочным обвинением.")
	fmt.Println()
	fmt.Println("Позже становится ясно:")
	fmt.Println("настоящий убийца остался на свободе.")
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println(Red + "              ДЕЛО ПРОВАЛЕНО" + Reset)
	fmt.Println("========================================")
	fmt.Println()
}
