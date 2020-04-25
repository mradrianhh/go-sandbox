package story

import "fmt"

// StartChapter2 ...
func StartChapter2(player *Player) {
	for {
		fmt.Print("\nChapter 2\n")
		fmt.Print("\n\t*You spot a virgin in need! She's being raped by a group of men!*\n")
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
}
