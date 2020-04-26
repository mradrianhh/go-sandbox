package game

import (
	"fmt"
)

// ShowChooseCharacterMenu ...
func ShowChooseCharacterMenu() {
	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Print("\n\nChoose Player")
		fmt.Printf("\n\n1 - %s. 0 - Main Menu\n", Player.Name)
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				fmt.Printf("\nWelcome to our adventure, %s!\n", Player.Name)
				switch Player.CurrentChapter {
				case 0:
					StartChapter0(&Player)
				case 1:
					StartChapter1(&Player)
				case 2:
					StartChapter2(&Player)
				case 3:
					StartChapter3(&Player)
				default:
					StartChapter0(&Player)
				}
			case 0:
				ShowMainMenu()
			default:
				fmt.Println("Sorry, I don't understand. Please try again.")
			}
		}
	}
}
