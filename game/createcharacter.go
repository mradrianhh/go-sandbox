package game

import (
	"fmt"
)

// ShowCreateCharacterMenu ...
func ShowCreateCharacterMenu() {
	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Println("1 - Name. 2 - Hair. 3 - Height. 4 - Race. 5 - Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				fmt.Print("\n----------------------------------------------------------\n")
				fmt.Println("\nChoose a Name.")
				fmt.Print("Enter your Name: ")
				if _, err := fmt.Scan(&Player.Name); err == nil {
					fmt.Printf("Welcome %s!\n", Player.Name)
					goto exitname
				}
			exitname:
				break
			case 2:
				for {
					fmt.Print("\n----------------------------------------------------------\n")
					fmt.Println("\nDesign your Hair.")
					fmt.Print("\nEnter your desired Hair color: ")
					if _, err := fmt.Scan(&Player.Hair); err == nil {
						fmt.Printf("Aah, I salute you %s of %s Hair.\n\n", Player.Name, Player.Hair)
						goto exithair
					}
				}
			exithair:
				break
			case 3:
				for {
					fmt.Print("\n----------------------------------------------------------\n")
					fmt.Println("\nChoose your Height from the menu below.")
					fmt.Println("\n1 - Tall. 2 - Short.")
					if _, err := fmt.Scan(&numResponse); err == nil {
						switch numResponse {
						case 1:
							Player.Height = "tall"
							goto exitheight
						case 2:
							Player.Height = "small"
							goto exitheight
						default:
							fmt.Println("Sorry, I don't understand. Please try again.")
						}
					}
				}
			exitheight:
				fmt.Printf("Aah, I salute you %s %s of %s Hair.\n\n", Player.Height, Player.Name, Player.Hair)
				break
			case 4:
				for {
					fmt.Print("\n----------------------------------------------------------\n")
					fmt.Println("\nChoose your Race.")
					fmt.Println("1 - Elf. 2 - Sorcerer.")
					if _, err := fmt.Scan(&numResponse); err == nil {
						switch numResponse {
						case 1:
							Player.Race = "elf"
							goto exitrace
						case 2:
							Player.Race = "sorcerer"
							goto exitrace
						default:
							fmt.Println("Sorry, I don't understand. Please try again.")
						}
					}
				}
			exitrace:
				fmt.Printf("Hmm... So %s %s of %s Hair wants to be an %s? So be it.\n\n", Player.Height, Player.Name, Player.Hair, Player.Race)
				break
			case 5:
				ShowMainMenu()
			default:
				fmt.Println("Sorry, I don't understand. Please try again.")
			}
		}
	}
}
