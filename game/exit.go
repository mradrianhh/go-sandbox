package game

import (
	"fmt"
	"os"
)

// ExitGame ...
func ExitGame() {
	fmt.Print("\n----------------------------------------------------------\n")
	fmt.Println("Goodbye!")
	os.Exit(0)
}
