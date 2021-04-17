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

	ch := make(chan string)
	go down(ch)
	go eventHandler(ch)
	for {
		_, key, _ := keyboard.GetKey()
		switch key {
		case keyboard.KeyArrowRight:
			ch <- "right"
		case keyboard.KeyArrowUp:
			ch <- "up"
		case keyboard.KeyArrowLeft:
			ch <- "left"
		case keyboard.KeyArrowDown:
			ch <- "down"
		case keyboard.KeyEsc:
			return
		}
	}

	tool.Print(arr)
}

func eventHandler(ch chan string) {
	var input = <- ch
	fmt.Println(input)
	go eventHandler(ch)
}

func down(ch chan string) {
	time.Sleep(time.Duration(1) * time.Second)
	ch <- "down"
	go down(ch)
}