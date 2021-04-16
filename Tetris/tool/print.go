package tool

import (
	"fmt"
	"github.com/fatih/color"
)


//yellow := color.New(color.FgYellow).SprintFunc()
//fmt.Printf("%s", yellow("  "))
//

//fmt.Printf("This %s rocks!\n", info("package"))
func Print(arr *[23][10]int) {
	println()
	println()
	border := color.New(color.BgHiWhite).SprintFunc()
	for i := 0; i < 24; i++ {
		switch i {
		case 0:
			fmt.Printf("      %s                    %s", border("  "), border("  "))
			fmt.Println()
			continue
		case 1:
			fmt.Printf("      %s                    %s", border("  "), border("  "))
			fmt.Println()
			continue
		case 2:
			fmt.Printf("      %s﹘﹘﹘﹘﹘﹘﹘﹘﹘﹘%s", border("  "), border("  "))
			fmt.Println()
			continue
		case 3:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {
					colorPrint(arr[i-3][j])
				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 4:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 5:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 6:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 7:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 8:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 9:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 11:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 12:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 13:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 14:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 15:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 16:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 17:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 18:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 19:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 20:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 21:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 22:
			fmt.Printf("      %s", border("  "))
			for j := 0; j < 10; j++ {
				if arr[i-3][j] == 0 {
					print("  ")
				} else {

				}
			}
			fmt.Printf("%s", border("  "))
			fmt.Println()
			continue
		case 23:
			fmt.Printf("      %s%s%s%s%s%s%s%s%s%s%s%s", border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "), border("  "))
			fmt.Println()
			continue
		}
	}
}

func colorPrint(a int) {
	color.New(color.BgHiCyan).SprintFunc()
	color.New(color.BgRed).SprintFunc()
	color.New(color.BgMagenta).SprintFunc()
	color.New(color.BgHiBlue).SprintFunc()
	color.New(color.BgBlue).SprintFunc()
	color.New(color.BgGreen).SprintFunc()
	color.New(color.BgYellow).SprintFunc()
	//fmt.Printf("%s", v("  "))
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