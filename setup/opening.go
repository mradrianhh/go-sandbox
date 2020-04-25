package setup

import "fmt"

// ShowOpening ...
func ShowOpening() {
	for {
		fmt.Print("Do you wanna play a game?(Y/N): ")
		if _, err := fmt.Scan(&charResponse); err == nil {
			switch charResponse {
			case "y", "Y":
				fmt.Println("Wonderful! Let's begin then.")
				goto mainmenu
			case "n", "N":
				fmt.Println("Hmmm... Fine then.")
				goto end
			}
		}
	}
}
