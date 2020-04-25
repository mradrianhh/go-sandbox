package main

import "fmt"

// Player ...
type Player struct {
	name           string
	hair           string
	height         string
	race           string
	currentChapter int
	empathy        int
	violence       int
	intelligence   int
}

func main() {
	var player Player
	player.name = "tbd"
	player.height = "tbd"
	player.hair = "tbd"
	player.race = "tbd"
	player.currentChapter = 0
	player.empathy = 10
	player.violence = 10
	player.intelligence = 10

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
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Print("\nMain Menu\n")
		fmt.Println("\n1 - Start Game. 2 - View Character. 3 - Create/Edit Character. 4. Exit")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				goto choosecharacter
			case 2:
				goto viewcharacter
			case 3:
				goto createcharacter
			case 4:
				goto end
			default:
				fmt.Println("Sorry, I don't understand. Please try again.")
			}
		}
	}

createcharacter:
	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Println("\nPlease, create your character if you will.")
		fmt.Print("Enter your name: ")
		if _, err := fmt.Scan(&player.name); err == nil {
			fmt.Printf("Welcome %s! Further design your character from the menu below.\n", player.name)
			goto designcharacter
		}
	}

designcharacter:
	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Println("1 - Hair. 2 - Height. 3 - Race. 4 - Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
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
			case 2:
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
			case 3:
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
			case 4:
				goto mainmenu
			default:
				fmt.Println("Sorry, I don't understand. Please try again.")
			}
		}
	}

viewcharacter:
	fmt.Print("\n----------------------------------------------------------\n")
	fmt.Print("\nView Character\n\n")
	fmt.Printf("\tName: %s\n", player.name)
	fmt.Printf("\tRace: %s\n", player.race)
	fmt.Printf("\tHeight: %s\n", player.height)
	fmt.Printf("\tHair: %s\n", player.hair)
	goto mainmenu

choosecharacter:
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

chapter0:
	for {
		fmt.Print("\nChapter 0\n")
		fmt.Print("\nThere was once a time where the world was good.\n")
		fmt.Print("Can you remember it, son?(Y/N)")
		if _, err := fmt.Scan(&charResponse); err == nil {
			switch charResponse {
			case "y", "Y":
				fmt.Println("Really? You were so young. Anyways...")
			case "n", "N":
				fmt.Println("That's fine, I don't expect you to. You were so young. Anyways...")
			default:
				fmt.Println("Really? You were so young. Anyways...")
			}
		}
		fmt.Print("Along the way, something happened. I don't know what. I don't know who. I just know... I just know that something happened.\n")
		fmt.Print("The world isn't supposed to be like this. So careless, so ruthless - so loveless...\n")
		fmt.Print("I can see so many humans, yet so little humanity. Like a child without joy, a child without sadness - an emotionless child.\n")
		fmt.Print("Promise me, son, to always strive to make the world a better place. Let the world know that we, as humans, are here. Make the world humane again!\n")

		fmt.Println("\n1 - Continue. 2 - Repeat. 3 - Quit to Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				goto chapter1
			case 2:
				goto chapter0
			case 3:
				goto mainmenu
			default:
				goto chapter1
			}
		}
	}

chapter1:
	for {
		fmt.Print("\nChapter 1\n")
		fmt.Print("\n\t*You spot a man in need*\n")
		fmt.Println("\n1 - Help him. 2 - Decapitate him. 3 - Yeet.")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				fmt.Println("You're just like your father. He trained you well. Good job!")
				player.empathy++
				player.violence--
				player.intelligence--
				fmt.Println("\n\tEmpathy increased by 1.")
				fmt.Println("\tViolence decreased by 1.")
				fmt.Println("\tIntelligence decreased by 1.")
			case 2:
				fmt.Println("Hahaha, just like I would've done it. Good to see you're father didn't leave any marks!")
				player.empathy--
				player.violence++
				player.intelligence++
				fmt.Println("\n\tEmpathy decreased by 1.")
				fmt.Println("\tViolence increased by 1.")
				fmt.Println("\tIntelligence increased by 1.")
			case 3:
				fmt.Println("You fucking coward, get back here! You'll regret this!")
				player.empathy--
				player.violence--
				player.intelligence--
				fmt.Println("\n\tEmpathy decreased by 1.")
				fmt.Println("\tViolence decreased by 1.")
				fmt.Println("\tIntelligence decreased by 1.")
			default:
				fmt.Println("You're just like your father. He trained you well. Good job!")
				player.empathy++
				player.violence--
				player.intelligence--
				fmt.Println("\n\tEmpathy increased by 1.")
				fmt.Println("\tViolence decreased by 1.")
				fmt.Println("\tIntelligence decreased by 1.")
			}
		}

		fmt.Println("\n1 - Continue. 2 - Repeat. 3 - Quit to Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				goto chapter2
			case 2:
				goto chapter1
			case 3:
				goto mainmenu
			default:
				goto chapter2
			}
		}
	}

chapter2:
	for {
		fmt.Print("\nChapter 2\n")
		fmt.Print("\n\t*You spot a virgin in need!*\n")
		fmt.Println("\n1 - Kill them. 2 - Rape her. 3 - Rape them. 4 - Yeet.")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				fmt.Println("\nWonderful! Perfect! I keep searching for a word to describe my joy over your decision, to no success.")
				player.empathy++
				player.violence++
				player.intelligence++
				fmt.Println("\n\tEmpathy increased by 1.")
				fmt.Println("\tViolence increased by 1.")
				fmt.Println("\tIntelligence increased by 1.")
			case 2:
				fmt.Println("\nHah, can't blame you, she was pretty hot, wasn't she? Well, moving on...")
				player.empathy--
				player.violence++
				player.intelligence++
				fmt.Println("\n\tEmpathy decreased by 1.")
				fmt.Println("\tViolence increased by 1.")
				fmt.Println("\tIntelligence increased by 1.")
			case 3:
				fmt.Println("\nUhm... So, care to explain? You know what? Don't. Let's move on...")
				player.empathy++
				player.empathy--
				player.violence++
				player.intelligence--
				fmt.Println("\n\tEmpathy increased by 1.")
				fmt.Println("\tEmpathy also decreased by 1.")
				fmt.Println("\tViolence increased by 1.")
				fmt.Println("\tIntelligence decreased by 1.")
			case 4:
				fmt.Println("\nOh for fucks sake! Come on! Can you stop it?")
				player.empathy--
				player.violence--
				player.intelligence--
				fmt.Println("\n\tEmpathy decreased by 1.")
				fmt.Println("\tViolence decreased by 1.")
				fmt.Println("\tIntelligence decreased by 1.")
			default:
				fmt.Println("\nWonderful! Perfect! I keep searching for a word to describe my joy over your decision, to no success.")
				player.empathy++
				player.violence++
				player.intelligence++
				fmt.Println("\n\tEmpathy increased by 1.")
				fmt.Println("\tViolence increased by 1.")
				fmt.Println("\tIntelligence increased by 1.")
			}
		}

		fmt.Println("\n1 - Continue. 2 - Repeat. 3 - Quit to Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				goto chapter3
			case 2:
				goto chapter2
			case 3:
				goto mainmenu
			default:
				goto chapter3
			}
		}
	}

chapter3:
	for {
		fmt.Print("\nChapter 3\n")
		fmt.Println("\nNow that we're back on track, let's see how you're doing so far.")

		if player.empathy > 5 {
			fmt.Printf("\n\tYour empathy is at %d, pretty decent, yeah?", player.empathy)
		} else {
			fmt.Printf("\n\tYour empathy is at %d, could've been worse?", player.empathy)
		}

		if player.violence > 5 {
			fmt.Printf("\n\tYour violence is at %d, pretty decent, yeah?", player.violence)
		} else {
			fmt.Printf("\n\tYour violence is at %d, could've been worse?", player.violence)
		}

		if player.intelligence > 5 {
			fmt.Printf("\n\tYour intelligence is at %d, pretty decent, yeah?", player.intelligence)
		} else {
			fmt.Printf("\n\tYour intelligence is at %d, could've been worse?", player.intelligence)
		}

		fmt.Println("\n\nDo keep in mind, if any of your stats reach 0, you lose. Good luck with the rest of your journey!")

		fmt.Println("\n1 - Continue. 2 - Repeat. 3 - Quit to Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				goto chapter4
			case 2:
				goto chapter2
			case 3:
				goto mainmenu
			default:
				goto chapter4
			}
		}
	}

chapter4:
	for {
		fmt.Print("\nChapter 4\n")

		fmt.Println("\n1 - Continue. 2 - Repeat. 3 - Quit to Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				goto chapter4
			case 2:
				goto chapter2
			case 3:
				goto mainmenu
			default:
				goto chapter4
			}
		}
	}

end:
	fmt.Print("\n----------------------------------------------------------\n")
	fmt.Println("Goodbye!")
}
