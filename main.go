package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Suspect struct {
	Name  string
	Age   int
	About string
	Alibi string
}

func ChooseSus(sus_col []Suspect, scanner *bufio.Scanner) {
	for {

		// текст выбрать подозреваемого
		fmt.Println()
		fmt.Println("Подозреваемые:")
		for i, sus := range sus_col {
			fmt.Printf("%d. %s, %d\n", i+1, sus.Name, sus.Age)
		}
		fmt.Println()
		fmt.Println("0. Назад")
		fmt.Println()
		fmt.Println("Выберите подозреваемого:")

		// получить число
		num_sus, ok := ReadNumber(scanner)
		if !ok {
			continue
		}

		// показать инфу
		if num_sus > 0 && num_sus <= len(sus_col) {
			fmt.Println()
			fmt.Printf("%s, %d\n", sus_col[num_sus-1].Name, sus_col[num_sus-1].Age)
			fmt.Println()
			fmt.Println("Описание:")
			fmt.Println(sus_col[num_sus-1].About)
			fmt.Println()
			fmt.Println("Алиби:")
			fmt.Println(sus_col[num_sus-1].Alibi)
			fmt.Println()
			return
		} else if num_sus == 0 {
			return
		} else {
			continue
		}
	}
}

func ReadNumber(scanner *bufio.Scanner) (int, bool) {
	valid := true

	if ok := scanner.Scan(); !ok {
		fmt.Println("Ошибка ввода!")
		valid = false
	}

	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("Пустой ввод!")
		valid = false
	}

	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Ошибка преобразования!", err)
		valid = false
	}

	return num, valid
}

func main() {
	susCollect := []Suspect{
		{
			Name:  "Joseph",
			Age:   54,
			About: "Строитель, необщительный, коренастый",
			Alibi: "Был на работе",
		},
		{
			Name:  "Monica",
			Age:   36,
			About: "Повар, качает права, низкая",
			Alibi: "Была с семьей",
		},
		{
			Name:  "Thomas",
			Age:   19,
			About: "Студент, относится несерьезно, высокий",
			Alibi: "Был на экзамене",
		},
		{
			Name:  "Karen",
			Age:   49,
			About: "Бухгалтер, хочет скорее уйти, среднее телосложение",
			Alibi: "Была на выставке",
		},
		{
			Name:  "Lilian",
			Age:   72,
			About: "На пенсии, за правосудие, маленькая",
			Alibi: "Была на рынке",
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {

		// Текст меню
		fmt.Println("ДЕЛО №17")
		fmt.Println("")
		fmt.Println("1. Исследовать место")
		fmt.Println("2. Допросить подозреваемого")
		fmt.Println("3. Посмотреть улики")
		fmt.Println("4. Обвинить")
		fmt.Println("5. Выйти")
		fmt.Println("")

		num, ok := ReadNumber(scanner)
		if !ok {
			continue
		}

		switch num {
		case 1:
			fmt.Println()
			fmt.Println("Выбрано 1 - Исследовать место")
			fmt.Println()

		case 2:

			ChooseSus(susCollect, scanner)

		case 3:
			fmt.Println()
			fmt.Println("Выбрано 3 - Посмотреть улики")
			fmt.Println()
		case 4:
			fmt.Println()
			fmt.Println("Выбрано 4 - Обвинить")
			fmt.Println()
		case 5:
			fmt.Println()
			fmt.Println("Выбрано 5 - Выйти")
			fmt.Println()
			return
		default:
			fmt.Println()
			fmt.Println("Неизвестная команда")
			fmt.Println()
		}
	}
}
