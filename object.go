package main

import (
	"bufio"
	"fmt"
)

type Object struct {
	LocID    int
	Name     string
	About    string
	Searched bool
	ObjID    int
	IsClue   bool
	Key      bool
}

func CreateObjects() []Object {
	objCollect := []Object{

		// ГОСТИНАЯ
		{
			ObjID:    1,
			LocID:    1,
			Name:     "🛋️   Диван",
			About:    "На диване лежит рюкзак Томаса.\nВнутри находится старая книга из коллекции Эдварда.",
			Searched: false,
			IsClue:   true,
		},
		{
			ObjID:    2,
			LocID:    1,
			Name:     "📚  Учебники",
			About:    "Редкие книги по истории и несколько листов с конспектами.\nСудя по записям, Томас готовился с ними к экзамену.",
			Searched: false,
			IsClue:   false,
		},
		{
			ObjID:    3,
			LocID:    1,
			Name:     "📖  Книга на журнальном столике",
			About:    "Между страницами книги спрятан сложенный вдвое лист бумаги.\nЭто личное письмо, адресованное Эмили.",
			Searched: false,
			IsClue:   true,
		},
		{
			ObjID:    4,
			LocID:    1,
			Name:     "☕️  Чай",
			About:    "Чай давно остыл.\nНичего необычного.",
			Searched: false,
			IsClue:   false,
		},

		// КУХНЯ
		{
			ObjID:    5,
			LocID:    2,
			Name:     "🕰️   Часы",
			About:    "Кухонные часы показывают время на 11 минут больше, чем ваши часы.\nПохоже, они давно спешат.",
			Searched: false,
			IsClue:   true,
		},
		{
			ObjID:    6,
			LocID:    2,
			Name:     "🗑️  Мусорное ведро",
			About:    "Обычный кухонный мусор: упаковки от продуктов, салфетки и пустая бутылка.\nНичего необычного.",
			Searched: false,
			IsClue:   false,
		},
		{
			ObjID:    7,
			LocID:    2,
			Name:     "🗄️  Нижний шкаф",
			About:    "Дверца одного из шкафов немного перекошена.\nЗа посудой ничего необычного нет, но задняя деревянная панель закреплена только с одной стороны.\nЗа панелью спрятана книга в тёмно-зелёном переплёте: «The Black Orchard».",
			Searched: false,
			IsClue:   true,
		},
		{
			ObjID:    8,
			LocID:    2,
			Name:     "🚪  Задняя дверь",
			About:    "Дверь ведёт во двор.\nЗамок исправен.\nСнаружи земля мокрая после дождя, но возле порога множество старых следов обуви — \nопределить, кому они принадлежат, невозможно.",
			Searched: false,
			IsClue:   false,
		},

		// БИБЛИОТЕКА
		{
			ObjID:    9,
			LocID:    3,
			Name:     "🩸  Тело",
			About:    "Эдвард лежит возле письменного стола.\nНа затылке глубокая рана.\nРядом на ковре лежит тяжёлый бронзовый подсвечник со следами крови.\nНа руке Эдварда разбиты наручные часы.\nСтрелки остановились на 20:37.",
			Searched: false,
			IsClue:   true,
		},
		{
			ObjID:    10,
			LocID:    3,
			Name:     "🖌️  Письменный стол",
			About:    "На столе лежат раскрытая записная книжка, несколько счетов и ручка.\nНа деревянном краю заметно небольшое пятно синей краски.",
			Searched: false,
			IsClue:   false,
		},
		{
			ObjID:    11,
			LocID:    3,
			Name:     "🪟  Окно",
			About:    "Окно закрыто изнутри.\nНа подоконнике лежит ровный слой пыли —\nследов рук или попытки открыть окно нет.",
			Searched: false,
			IsClue:   false,
		},
		{
			ObjID:    12,
			LocID:    3,
			Name:     "🕯️  Камин",
			About:    "В камине среди золы виден кусок бумаги, который не сгорел полностью.\nНа сохранившейся части можно разобрать:\nMORRIS CONSULTING\nInvoice #0417\n£4,800",
			Searched: false,
			IsClue:   true,
		},
		{
			ObjID:    13,
			LocID:    3,
			Name:     "📥  Витрина",
			About:    "Стеклянная витрина не повреждена, но одна полка пуста.\nНа табличке указано: «The Black Orchard. First edition».\nЗамок цел. Следов взлома нет.",
			Searched: false,
			IsClue:   false,
			Key:      true,
		},
	}
	return objCollect
}

func ChooseObjects(game *GameState, scanner *bufio.Scanner, choiceLoc int) {
	for {
		ClearScreen()

		// текст выбрать объект
		fmt.Println()
		fmt.Println("Предметы в комнате:")
		count := 0
		indexes := []int{}
		for i, obj := range game.Objects {
			if obj.LocID == choiceLoc {
				indexes = append(indexes, i)
				count += 1
				fmt.Printf("%d. %s\n", count, obj.Name)
			}
		}
		fmt.Println()
		fmt.Println("0. Назад")
		fmt.Println()
		fmt.Println("Выберите предмет для изучения:")

		choiceNum, ok := ReadNumber(scanner)
		if !ok {
			fmt.Println("Неверный ввод")
			continue
		}

		if choiceNum == 0 {
			ClearScreen()
			return
		} else if choiceNum < 1 || choiceNum > len(indexes) {
			ClearScreen()
			fmt.Println("Такого предмета нет.")
			continue
		} else {
			ClearScreen()
			// подогнать выбранный номер под комнату
			realIndex := indexes[choiceNum-1]
			ShowObject(game, realIndex, scanner)
			continue
		}
	}
}

func ShowObject(game *GameState, realIndex int, scanner *bufio.Scanner) {
	for {

		// показать инфу
		// не осматривали, найдена улика
		if game.Objects[realIndex].Searched == false &&
			game.Objects[realIndex].IsClue {
			game.Objects[realIndex].Searched = true
			fmt.Println()
			fmt.Println("====================")
			fmt.Printf("%s:\n", game.Objects[realIndex].Name)
			fmt.Println("====================")
			fmt.Println()
			fmt.Println(game.Objects[realIndex].About)

			FoundClue(game, game.Objects[realIndex].ObjID)

			Pause(scanner)
			ClearScreen()
			return

			// fmt.Println()
			// fmt.Println("0. Назад")
			// fmt.Println()

			// не осматривали, нет улики
		} else if game.Objects[realIndex].Searched == false &&
			game.Objects[realIndex].IsClue == false {
			game.Objects[realIndex].Searched = true
			fmt.Println()
			fmt.Println("====================")
			fmt.Printf("%s:\n", game.Objects[realIndex].Name)
			fmt.Println("====================")
			fmt.Println()
			fmt.Println(game.Objects[realIndex].About)

			Pause(scanner)
			ClearScreen()
			return

			// fmt.Println()
			// fmt.Println("0. Назад")
			// fmt.Println()

			// осматривали, была улика
		} else if game.Objects[realIndex].Searched == true &&
			game.Objects[realIndex].IsClue == true {
			fmt.Println()
			fmt.Println("====================")
			fmt.Printf("%s:\n", game.Objects[realIndex].Name)
			fmt.Println("====================")
			fmt.Println()
			fmt.Println("Вы уже осматривали это место.\nВсе важные находки отсюда уже добавлены в список улик.")

			Pause(scanner)
			ClearScreen()
			return

			// fmt.Println()
			// fmt.Println("0. Назад")
			// fmt.Println()

			// осматривали, не было улики
		} else if game.Objects[realIndex].Searched == true &&
			game.Objects[realIndex].IsClue == false {
			fmt.Println()
			fmt.Println("====================")
			fmt.Printf("%s:\n", game.Objects[realIndex].Name)
			fmt.Println("====================")
			fmt.Println()
			fmt.Println("Вы внимательно осматривали это место раньше.\nНичего нового обнаружить не удалось.")

			Pause(scanner)
			ClearScreen()
			return

			// fmt.Println()
			// fmt.Println("0. Назад")
			// fmt.Println()
		}
		// получить число
		// choiceNum, ok := ReadNumber(scanner)
		// if !ok {
		// 	fmt.Println("Неверный ввод")
		// 	continue
		// }

		// if choiceNum == 0 {
		// 	return
		// } else {
		// 	continue
		// }
	}
}

func KeyObj(game *GameState) int {
	keys := 0
	for _, obj := range game.Objects {
		if obj.Key && obj.Searched {
			keys += 1
		}
	}
	return keys
}
