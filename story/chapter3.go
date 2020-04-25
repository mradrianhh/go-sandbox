package story

import "fmt"

// StartChapter3 ...
func StartChapter3(player *Player) {
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
}
