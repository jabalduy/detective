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
	keys := KeyClues(game) + KeyDial(game) + KeyObj(game)

	if keys == 7 && choiceAc == 4 {
		fmt.Println(WinEnding())
		Pause(scanner)
	} else if keys > 3 && choiceAc == 4 {
		fmt.Println(UnsolvedEnding())
		Pause(scanner)
	} else {
		fmt.Println(FailedEnding())
		Pause(scanner)
	}
}

func WinEnding() string {
	return `========================================
              ОБВИНЕНИЕ
========================================

Вы называете имя:

Карен Моррис.

Следствие складывается в единую картину.

Эдвард погиб около 20:37.
В камине найден документ MORRIS CONSULTING,
связанный с финансовыми махинациями.

Карен утверждала, что приехала позже
и вообще не заходила в библиотеку.

Но её видели возле кухни раньше, чем она призналась,
а пропавшая книга была спрятана именно там.

Кража была инсценировкой.
Карен попыталась скрыть настоящий мотив убийства.

Несколько секунд она молчит.

Затем тихо говорит:

«Он обнаружил всё.
На следующее утро собирался обратиться к адвокату.
Я хотела забрать документы.
Мы поссорились...
Я не собиралась его убивать.»

========================================
              ДЕЛО ЗАКРЫТО
========================================`
}

func UnsolvedEnding() string {
	return `========================================
              ОБВИНЕНИЕ
========================================

Вы обвиняете Карен Моррис.

Версия выглядит правдоподобно,
но в расследовании остаются пробелы.

Часть важных улик не найдена,
некоторые показания не проверены,
а цепочка событий не доказана полностью.

Карен спокойно отрицает обвинение.

Полиции недостаточно оснований,
чтобы предъявить ей обвинение.

Карен покидает дом.

Вы знаете, что могли быть правы.
Но знать — недостаточно.

========================================
             НЕДОСТАТОЧНО УЛИК
========================================`
}

func FailedEnding() string {
	return `========================================
              ОБВИНЕНИЕ
========================================

Вы делаете окончательный выбор.

Но собранные факты не складываются
в доказательство вины этого человека.

Расследование заканчивается ошибочным обвинением.

Позже становится ясно:
настоящий убийца остался на свободе.

========================================
              ДЕЛО ПРОВАЛЕНО
========================================`
}
