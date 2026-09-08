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
			Name:  "💌  Письмо Эмили",
			About: "Личное письмо Томаса к Эмили.\nОно подтверждает, что Томас скрывал их отношения\nи планировал встретиться с ней вечером.",
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
		},
		{
			Name:  "⌚️  Разбитые часы",
			About: "Наручные часы Эдварда разбиты при падении.\nСтрелки остановились на 20:37.",
			Found: false,
			ObjID: 9,
		},
		{
			Name:  "📜  Обгоревший документ",
			About: "Фрагмент финансового документа, найденный в камине библиотеки.\nНа нём сохранились надписи\n«MORRIS CONSULTING»,\n«Invoice #0417»\nи сумма £4,800.",
			Found: false,
			ObjID: 12,
		},
	}
	return clueCollect
}

func ShowClues(game *GameState, scanner *bufio.Scanner) {
	for {

		// показать найденные улики
		fmt.Println()
		fmt.Println("=========")
		fmt.Println("Улики:")
		fmt.Println("=========")
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
	fmt.Println(`
        \   |   /
      '.  \ | /  .'
    ---   * * *   ---
      .'  / | \  '.
        /   |   \

      НАЙДЕНА УЛИКА
	  `)
}
