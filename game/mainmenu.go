package game

import "fmt"

var numResponse int
var charResponse string

// ShowMainMenu ...
func ShowMainMenu() {
	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Print("\nMain Menu\n")
		fmt.Println("\n1 - Start Game. 2 - View Character. 3 - Create/Edit Character. 4. Exit")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				ShowChooseCharacterMenu()
			case 2:
				ViewCharacter()
			case 3:
				ShowCreateCharacterMenu()
			case 4:
				ExitGame()
			default:
				fmt.Println("Sorry, I don't understand. Please try again.")
			}
		}
	}
}
