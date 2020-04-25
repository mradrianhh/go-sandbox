package story

import "fmt"

func StartChapter1(player *Player) {
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
}
