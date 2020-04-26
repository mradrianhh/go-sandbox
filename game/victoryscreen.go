package game

import "fmt"

// ShowVictoryScreen ...
func ShowVictoryScreen() {
	fmt.Print("\n----------------------------------------------------------\n")
	fmt.Print("\n*************************You Won**************************\n")
	ShowMainMenu()
}
