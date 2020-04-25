package setup

import "fmt"

// ShowCreateCharacterMenu ...
func ShowCreateCharacterMenu(player *Player) {
	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Println("1 - Name. 2 - Hair. 3 - Height. 4 - Race. 5 - Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				fmt.Print("\n----------------------------------------------------------\n")
				fmt.Println("\nChoose a name.")
				fmt.Print("Enter your name: ")
				if _, err := fmt.Scan(&player.name); err == nil {
					fmt.Printf("Welcome %s!\n", player.name)
					goto exitname
				}
			exitname:
				break
			case 2:
				for {
					fmt.Print("\n----------------------------------------------------------\n")
					fmt.Println("\nDesign your hair.")
					fmt.Print("\nEnter your desired hair color: ")
					if _, err := fmt.Scan(&player.hair); err == nil {
						fmt.Printf("Aah, I salute you %s of %s hair.\n\n", player.name, player.hair)
						goto exithair
					}
				}
			exithair:
				break
			case 3:
				for {
					fmt.Print("\n----------------------------------------------------------\n")
					fmt.Println("\nChoose your height from the menu below.")
					fmt.Println("\n1 - Tall. 2 - Short.")
					if _, err := fmt.Scan(&numResponse); err == nil {
						switch numResponse {
						case 1:
							player.height = "tall"
							goto exitheight
						case 2:
							player.height = "small"
							goto exitheight
						default:
							fmt.Println("Sorry, I don't understand. Please try again.")
						}
					}
				}
			exitheight:
				fmt.Printf("Aah, I salute you %s %s of %s hair.\n\n", player.height, player.name, player.hair)
				break
			case 4:
				for {
					fmt.Print("\n----------------------------------------------------------\n")
					fmt.Println("\nChoose your race.")
					fmt.Println("1 - Elf. 2 - Sorcerer.")
					if _, err := fmt.Scan(&numResponse); err == nil {
						switch numResponse {
						case 1:
							player.race = "elf"
							goto exitrace
						case 2:
							player.race = "sorcerer"
							goto exitrace
						default:
							fmt.Println("Sorry, I don't understand. Please try again.")
						}
					}
				}
			exitrace:
				fmt.Printf("Hmm... So %s %s of %s hair wants to be an %s? So be it.\n\n", player.height, player.name, player.hair, player.race)
				break
			case 5:
				goto mainmenu
			default:
				fmt.Println("Sorry, I don't understand. Please try again.")
			}
		}
	}
}
