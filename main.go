package main

import "fmt"

// Player ...
type Player struct {
	name   string
	hair   string
	height string
	race   string
}

func main() {
	var player Player
	player.name = "tbd"
	player.height = "tbd"
	player.hair = "tbd"
	player.race = "tbd"

	var numResponse int
	var charResponse string

	for {
		fmt.Print("Do you wanna play a game?(Y/N): ")
		if _, err := fmt.Scan(&charResponse); err == nil {
			switch charResponse {
			case "y", "Y":
				fmt.Println("Wonderful! Let's begin then.")
				goto mainmenu
			case "n", "N":
				fmt.Println("Hmmm... Fine then.")
				goto end
			}
		}
	}

mainmenu:
	for {
		fmt.Println("\n1 - Start Game. 2 - View Character. 3 - Create/Edit Character. 4. Exit")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				goto startgame
			case 2:
				goto viewcharacter
			case 3:
				goto createcharacter
			case 4:
				fmt.Println("Goodbye!")
				goto end
			default:
				fmt.Println("Sorry, I don't understand. Please try again.")
			}
		}
	}

createcharacter:
	for {
		fmt.Println("\nPlease, create your character if you will.")
		fmt.Print("Enter your name: ")
		if _, err := fmt.Scan(&player.name); err == nil {
			fmt.Printf("Welcome %s! Further design your character from the menu below.\n", player.name)
			goto designcharacter
		}
	}

designcharacter:
	for {
		fmt.Println("1 - Hair. 2 - Height. 3 - Race. 4 - Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				for {
					fmt.Println("Design your hair.")
					fmt.Print("Enter your desired hair color: ")
					if _, err := fmt.Scan(&player.hair); err == nil {
						fmt.Printf("Aah, I salute you %s of %s hair.\n\n", player.name, player.hair)
						goto exithair
					}
				}
			exithair:
				break
			case 2:
				for {
					fmt.Println("Choose your height from the menu below.")
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
			case 3:
				for {
					fmt.Println("Choose your race.")
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
			case 4:
				goto mainmenu
			default:
				fmt.Println("Sorry, I don't understand. Please try again.")
			}
		}
	}

viewcharacter:
	fmt.Print("View Character\n\n")
	fmt.Printf("\tName: %s\n", player.name)
	fmt.Printf("\tRace: %s\n", player.race)
	fmt.Printf("\tHeight: %s\n", player.height)
	fmt.Printf("\tHair: %s\n", player.hair)
	goto mainmenu

startgame:
	fmt.Printf("Welcome to our adventure, %s!\n", player.name)

end:
}
