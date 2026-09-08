package main

import (
	"bufio"
	"fmt"
	"os"
)

type GameState struct {
	Suspects  []Suspect
	Clues     []Clue
	Locations []Location
	Objects   []Object
	Dialogues []Dialogue
	End       bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	game := GameState{
		Suspects:  CreateSuspects(),
		Clues:     CreateClue(),
		Locations: CreateLocations(),
		Objects:   CreateObjects(),
		Dialogues: CreateDialogue(),
		End:       false,
	}

	PrintIntro()

	for {
		// Текст меню
		fmt.Println("")
		fmt.Println("ДЕЛО №17")
		fmt.Println("")
		fmt.Println("1. Исследовать комнаты")
		fmt.Println("2. Допросить подозреваемого")
		fmt.Println("3. Открыть Досье")
		fmt.Println("4. Обвинить")
		fmt.Println("5. Выйти")
		fmt.Println("")

		num, ok := ReadNumber(scanner)
		if !ok {
			continue
		}

		switch num {
		case 1:

			ChooseLocations(&game, scanner)

		case 2:

			ChooseSus(&game, scanner)

		case 3:

			ShowCaseFile(&game, scanner)

		case 4:

			if ChooseAccused(&game, scanner) {
				return
			} else {
				continue
			}

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
