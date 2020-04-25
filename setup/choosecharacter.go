package setup

import "fmt"

// ShowChooseCharacterMenu ...
func ShowChooseCharacterMenu() {
	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Print("\n\nChoose player")
		fmt.Printf("\n\n1 - %s. 0 - Main Menu\n", player.name)
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				fmt.Printf("Welcome to our adventure, %s!\n", player.name)
				switch player.currentChapter {
				case 0:
					goto chapter0
				case 1:
					goto chapter1
				case 2:
					goto chapter2
				default:
					goto chapter0
				}
			case 0:
				goto mainmenu
			default:
				fmt.Println("Sorry, I don't understand. Please try again.")
			}
		}
	}
}
