package models

// Player ...
type Player struct {
	Name           string
	Hair           string
	Height         string
	Race           string
	CurrentChapter int
	Empathy        int
	Violence       int
	Intelligence   int
}

// GetPlayer returns the saved characters in the game.
func GetPlayer() Player {
	var player1 Player
	player1.Name = "Adrian"
	player1.Height = "tall"
	player1.Hair = "black"
	player1.Race = "elf"
	player1.CurrentChapter = 0
	player1.Empathy = 5
	player1.Violence = 5
	player1.Intelligence = 5

	/* var player2 Player
	player2.Name = "Jack"
	player2.Height = "short"
	player2.Hair = "black"
	player2.Race = "elf"
	player2.CurrentChapter = 0
	player2.Empathy = 5
	player2.Violence = 5
	player2.Intelligence = 5 */

	//return []Player{player1, player2}

	return player1
}
