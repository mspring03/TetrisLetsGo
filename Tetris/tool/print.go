package tool

import (
	"fmt"
	"github.com/fatih/color"
	"os/exec"
)


//yellow := color.New(color.FgYellow).SprintFunc()
//fmt.Printf("%s", yellow("  "))
//

//fmt.Printf("This %s rocks!\n", info("package"))
func Print(arr [][]int) {
	output, _ := exec.Command("clear").CombinedOutput()
	fmt.Print(string(output))
	println()
	println()
	border := color.New(color.BgHiWhite).SprintFunc()
	for i := 22; i >= 0; i-- {
		switch i {
		case 20:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i][j] == 0 {
					print("﹘")
				} else {
					colorPrint(arr[i][j])
				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 0:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i][j] == 0 {
					print("  ")
				} else {
					colorPrint(arr[i][j])
				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			fmt.Printf("      %s%s%s%s%s%s%s%s%s%s%s%s", border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "))
			fmt.Println()
			continue
		default:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i][j] == 0 {
					print("  ")
				} else {
					colorPrint(arr[i][j])
				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		}
	}
}

func colorPrint(a int) {
	switch a {
	case 1:
		Cyan := color.New(color.BgHiCyan).SprintFunc()
		fmt.Printf("%s", Cyan("  "))
		break
	case 2:
		Red := color.New(color.BgRed).SprintFunc()
		fmt.Printf("%s", Red("  "))
		break
	case 3:
		Magenta := color.New(color.BgMagenta).SprintFunc()
		fmt.Printf("%s", Magenta("  "))
		break
	case 4:
		HiBlue := color.New(color.BgHiBlue).SprintFunc()
		fmt.Printf("%s", HiBlue("  "))
		break
	case 5:
		Green := color.New(color.BgGreen).SprintFunc()
		fmt.Printf("%s", Green("  "))
		break
	case 6:
		Yellow := color.New(color.BgYellow).SprintFunc()
		fmt.Printf("%s", Yellow("  "))
		break
	case 7:
	}
}