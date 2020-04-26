package game

import (
	"fmt"

	"github.com/mradrianhh/go-sandbox/models"
)

// Player ...
var Player models.Player

func init() {
	Player = models.GetPlayer()
}

// ViewCharacter ...
func ViewCharacter() {
	fmt.Print("\n----------------------------------------------------------\n")
	fmt.Print("\nView Character\n\n")
	fmt.Printf("\tName: %s\n", Player.Name)
	fmt.Printf("\tRace: %s\n", Player.Race)
	fmt.Printf("\tHeight: %s\n", Player.Height)
	fmt.Printf("\tHair: %s\n", Player.Hair)
	ShowMainMenu()
}
