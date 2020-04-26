package game

import (
	"fmt"

	"github.com/mradrianhh/go-sandbox/models"
)

// StartChapter1 ...
func StartChapter1(player *models.Player) {
	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Print("\nChapter 1\n")
		fmt.Print("\n\t*You spot a man in need*\n")
		fmt.Println("\n1 - Help him. 2 - Decapitate him. 3 - Yeet.")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				fmt.Println("You're just like your father. He trained you well. Good job!")
				player.Empathy++
				player.Violence--
				player.Intelligence--
				fmt.Println("\n\tEmpathy increased by 1.")
				fmt.Println("\tViolence decreased by 1.")
				fmt.Println("\tIntelligence decreased by 1.")
			case 2:
				fmt.Println("Hahaha, just like I would've done it. Good to see you're father didn't leave any marks!")
				player.Empathy--
				player.Violence++
				player.Intelligence++
				fmt.Println("\n\tEmpathy decreased by 1.")
				fmt.Println("\tViolence increased by 1.")
				fmt.Println("\tIntelligence increased by 1.")
			case 3:
				fmt.Println("You fucking coward, get back here! You'll regret this!")
				player.Empathy--
				player.Violence--
				player.Intelligence--
				fmt.Println("\n\tEmpathy decreased by 1.")
				fmt.Println("\tViolence decreased by 1.")
				fmt.Println("\tIntelligence decreased by 1.")
			default:
				fmt.Println("You're just like your father. He trained you well. Good job!")
				player.Empathy++
				player.Violence--
				player.Intelligence--
				fmt.Println("\n\tEmpathy increased by 1.")
				fmt.Println("\tViolence decreased by 1.")
				fmt.Println("\tIntelligence decreased by 1.")
			}
		}

		fmt.Println("\n1 - Continue. 2 - Repeat. 3 - Quit to Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				StartChapter2(player)
			case 2:
				StartChapter1(player)
			case 3:
				ShowMainMenu()
			default:
				StartChapter2(player)
			}
		}
	}
}
