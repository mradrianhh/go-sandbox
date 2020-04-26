package game

import (
	"fmt"

	"github.com/mradrianhh/go-sandbox/models"
)

// StartChapter3 ...
func StartChapter3(player *models.Player) {
	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Print("\nChapter 3\n")
		fmt.Println("\nNow that we're back on track, let's see how you're doing so far.")

		if player.Empathy > 5 {
			fmt.Printf("\n\tYour empathy is at %d, pretty decent, yeah?", player.Empathy)
		} else {
			fmt.Printf("\n\tYour empathy is at %d, could've been worse?", player.Empathy)
		}

		if player.Violence > 5 {
			fmt.Printf("\n\tYour violence is at %d, pretty decent, yeah?", player.Violence)
		} else {
			fmt.Printf("\n\tYour violence is at %d, could've been worse?", player.Violence)
		}

		if player.Intelligence > 5 {
			fmt.Printf("\n\tYour intelligence is at %d, pretty decent, yeah?", player.Intelligence)
		} else {
			fmt.Printf("\n\tYour intelligence is at %d, could've been worse?", player.Intelligence)
		}

		fmt.Println("\n\nDo keep in mind, if any of your stats reach 0, you lose. Good luck with the rest of your journey!")

		fmt.Println("\n1 - Continue. 2 - Repeat. 3 - Quit to Main Menu")
		if _, err := fmt.Scan(&numResponse); err == nil {
			switch numResponse {
			case 1:
				ShowVictoryScreen()
			case 2:
				StartChapter3(player)
			case 3:
				ShowMainMenu()
			default:
				ShowVictoryScreen()
			}
		}
	}
}
