package lib

import (
	"math"
	"os"
	"os/exec"
)

// Function to clear the terminal, test on Windows
func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CenterText(title string, ColLen int) (titleSpaceLeft, titleSpaceRight int) {
	titleLen = len(title)
	titleHalfLen = float64(ColLen-titleLen) / 2
	titleSpaceLeft = int(math.Ceil(titleHalfLen))
	titleSpaceRight = int(math.Floor(titleHalfLen))
	return
}
