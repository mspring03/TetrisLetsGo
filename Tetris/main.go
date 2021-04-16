package main

import (
	"fmt"
	"github.com/mspring03/TetrisLetsGo/Tetris/tool"
	"os/exec"
)

func main() {
	//for {
		output, _ := exec.Command("clear").CombinedOutput()
		fmt.Print(string(output))


	//}
	// Print with default helper functions
	//color.Cyan("■")
	var arr [23][10]int
	tool.Print(&arr)



}
