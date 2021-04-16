package main

import "fmt"

func Print() {
	println()
	println()
	for i := 0; i < 22; i++ {
		switch i {
		case 0:
			fmt.Print("   ▩")
			fmt.Print("_ _ _ _ _ ")
			fmt.Print("▩▩▩▩▩▩▩▩▩▩")
		}
	}
}