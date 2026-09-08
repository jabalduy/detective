package main

import (
	"bufio"
	"fmt"
)

type Clue struct {
	Name  string
	About string
	Found bool
	ObjID int
	Key   bool
}

func CreateClue() []Clue {
	clueCollect := []Clue{
		{
			Name:  "🪪  Карточка выдачи",
			About: "В книге из рюкзака Томаса обнаружена карточка:\n«Томасу - вернуть в пятницу».\nНиже стоит подпись Эдварда Вейла.",
			Found: false,
			ObjID: 1,
		},
		{
			Name: "💌  Письмо Эмили",
			About: "\n" +
				"«Эмили, я больше не могу делать вид, что всё нормально.\n" +
				"Твой дед не имеет права решать, с кем тебе быть.\n" +
				"\n" +
				"Если он ещё раз попытается нас разлучить, я сам с ним поговорю.\n" +
				"Мне всё равно, насколько ему это не понравится.\n" +
				"\n" +
				"Я устал от него и от того, что мы должны постоянно скрываться.\n" +
				"Томас»",
			Found: false,
			ObjID: 3,
		},
		{
			Name:  "🕰️   Кухонные часы",
			About: "Часы на кухне спешат ровно на 11 минут.",
			Found: false,
			ObjID: 5,
		},
		{
			Name:  "📙  Пропавшая книга",
			About: "Первое издание «The Black Orchard»,\nякобы украденное из библиотеки.\nКнига была спрятана за задней панелью кухонного шкафа и не покидала дом.",
			Found: false,
			ObjID: 7,
			Key:   true,
		},
		{
			Name:  "⌚️  Разбитые часы",
			About: "Наручные часы Эдварда разбиты при падении.\nСтрелки остановились на 20:37.",
			Found: false,
			ObjID: 9,
			Key:   true,
		},
		{
			Name:  "📜  Обгоревший документ",
			About: "Фрагмент финансового документа, найденный в камине библиотеки.\nНа нём сохранились надписи\n«MORRIS CONSULTING»,\n«Invoice #0417»\nи сумма £4,800.",
			Found: false,
			ObjID: 12,
			Key:   true,
		},
	}
	return clueCollect
}

func ShowClues(game *GameState, scanner *bufio.Scanner) {

	// показать найденные улики
	fmt.Println()
	fmt.Println("УЛИКИ")
	fmt.Println()
	count := 0
	for _, clue := range game.Clues {
		if clue.Found {
			count += 1
			fmt.Printf("%d. %s\n", count, clue.Name)
			fmt.Println(clue.About)
			fmt.Println()
		}
	}
	if count == 0 {
		fmt.Println("Улики пока не найдены.")
		fmt.Println()
	}
}

func FoundClue(game *GameState, objID int) {
	for index, clue := range game.Clues {
		if clue.ObjID == objID {
			PrintStar()
			fmt.Printf("%s:\n", clue.Name)
			fmt.Println(clue.About)
			game.Clues[index].Found = true
			for i, dial := range game.Dialogues {
				if dial.ObjID == clue.ObjID {
					game.Dialogues[i].IsClue = true
					game.Dialogues[i].IsOpen = true
					fmt.Println()
					fmt.Println("💭 Новые диалоги разблокированы!")
					fmt.Println()
				}
			}
		}
	}
}

func PrintStar() {
	fmt.Println(Gold + `
        \   |   /
      '.  \ | /  .'
    ---   * * *   ---
      .'  / | \  '.
        /   |   \

      НАЙДЕНА УЛИКА
	  ` + Reset)
}

func KeyClues(game *GameState) int {
	keys := 0
	for _, clue := range game.Clues {
		if clue.Key && clue.Found {
			keys += 1
		}
	}
	return keys
}
