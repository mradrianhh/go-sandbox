package setup

import "fmt"

func ViewCharacter(player *Player) {
	fmt.Print("\n----------------------------------------------------------\n")
	fmt.Print("\nView Character\n\n")
	fmt.Printf("\tName: %s\n", player.name)
	fmt.Printf("\tRace: %s\n", player.race)
	fmt.Printf("\tHeight: %s\n", player.height)
	fmt.Printf("\tHair: %s\n", player.hair)
	goto mainmenu
}
