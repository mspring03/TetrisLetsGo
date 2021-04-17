package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/mspring03/TetrisLetsGo/Tetris/tool"
	"time"

	//"time"
)

func main() {
	arr := make([][]int,23)
	for i:=0 ; i<23; i++{
		arr[i] = make([]int,10)
	}

	//for {
	//	randShape := rand.Intn(7)+1
	//	point := []int{5, 22}
	//
	//	for {
	//
	//	}
	//}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")

	ch1 := make(chan string)
	ch2 := make(chan string)
	go down(ch1, ch2)
	go eventHandler(ch1, ch2)
	for {
		_, key, _ := keyboard.GetKey()
		switch key {
		case keyboard.KeyArrowRight:
			ch1 <- "right"
		case keyboard.KeyArrowUp:
			ch1 <- "up"
		case keyboard.KeyArrowLeft:
			ch1 <- "left"
		case keyboard.KeyArrowDown:
			ch1 <- "down"
		case keyboard.KeyEsc:
			ch1 <- "esc"
			return
		}
	}

	tool.Print(arr)
}

func eventHandler(ch1 chan string, ch2 chan string) {
	var input = <- ch1
	if input == "esc" {
		return
	}

	go eventHandler(ch1, ch2)
}

func down(ch1 chan string, ch2 chan string) {
	time.Sleep(time.Duration(1) * time.Second)
	ch1 <- "down"
	go down(ch1, ch2)
}