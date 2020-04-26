package game

import (
	"fmt"

	"github.com/mradrianhh/go-sandbox/models"
)

// StartChapter2 ...
func StartChapter2(player *models.Player) {
	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Print("\nChapter 2\n")
		fmt.Print("\n\t*You spot a virgin in need! She's being raped by a group of men!*\n")
		fmt.Println("\n1 - Kill them. 2 - Rape her. 3 - Rape them. 4 - Yeet.")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				fmt.Println("\nWonderful! Perfect! I keep searching for a word to describe my joy over your decision, to no success.")
				player.Empathy++
				player.Violence++
				player.Intelligence++
				fmt.Println("\n\tEmpathy increased by 1.")
				fmt.Println("\tViolence increased by 1.")
				fmt.Println("\tIntelligence increased by 1.")
			case 2:
				fmt.Println("\nHah, can't blame you, she was pretty hot, wasn't she? Well, moving on...")
				player.Empathy--
				player.Violence++
				player.Intelligence++
				fmt.Println("\n\tEmpathy decreased by 1.")
				fmt.Println("\tViolence increased by 1.")
				fmt.Println("\tIntelligence increased by 1.")
			case 3:
				fmt.Println("\nUhm... So, care to explain? You know what? Don't. Let's move on...")
				player.Empathy++
				player.Empathy--
				player.Violence++
				player.Intelligence--
				fmt.Println("\n\tEmpathy increased by 1.")
				fmt.Println("\tEmpathy also decreased by 1.")
				fmt.Println("\tViolence increased by 1.")
				fmt.Println("\tIntelligence decreased by 1.")
			case 4:
				fmt.Println("\nOh for fucks sake! Come on! Can you stop it?")
				player.Empathy--
				player.Violence--
				player.Intelligence--
				fmt.Println("\n\tEmpathy decreased by 1.")
				fmt.Println("\tViolence decreased by 1.")
				fmt.Println("\tIntelligence decreased by 1.")
			default:
				fmt.Println("\nWonderful! Perfect! I keep searching for a word to describe my joy over your decision, to no success.")
				player.Empathy++
				player.Violence++
				player.Intelligence++
				fmt.Println("\n\tEmpathy increased by 1.")
				fmt.Println("\tViolence increased by 1.")
				fmt.Println("\tIntelligence increased by 1.")
			}
		}

		fmt.Println("\n1 - Continue. 2 - Repeat. 3 - Quit to Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				StartChapter3(player)
			case 2:
				StartChapter2(player)
			case 3:
				ShowMainMenu()
			default:
				StartChapter3(player)
			}
		}
	}
}
