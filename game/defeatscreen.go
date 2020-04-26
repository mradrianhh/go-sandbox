package game

import "fmt"

// ShowDefeatScreen ...
func ShowDefeatScreen() {
	fmt.Print("\n----------------------------------------------------------\n")
	fmt.Print("\n*************************You Lost**************************\n")
	ShowMainMenu()
}
