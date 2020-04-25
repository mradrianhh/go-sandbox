package main

import "fmt"

// Player ...
type Player struct {
	name           string
	hair           string
	height         string
	race           string
	currentChapter int
	empathy        int
	violence       int
	intelligence   int
}

func main() {
	var player Player
	player.name = "tbd"
	player.height = "tbd"
	player.hair = "tbd"
	player.race = "tbd"
	player.currentChapter = 0
	player.empathy = 5
	player.violence = 5
	player.intelligence = 5

end:
	fmt.Print("\n----------------------------------------------------------\n")
	fmt.Println("Goodbye!")
}
