package game

import (
	"fmt"

	"github.com/mradrianhh/go-sandbox/models"
)

// StartChapter0 starts the first chapter.
func StartChapter0(player *models.Player) {

	for {
		fmt.Print("\n----------------------------------------------------------\n")
		fmt.Print("\nChapter 0\n")
		fmt.Print("\nThere was once a time where the world was good.\n")
		fmt.Printf("Can you remember it, %s?(Y/N)", player.Name)
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
				StartChapter1(player)
			case 2:
				StartChapter0(player)
			case 3:
				ShowMainMenu()
			default:
				StartChapter1(player)
			}
		}
	}
}
