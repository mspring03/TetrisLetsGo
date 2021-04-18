package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/mspring03/TetrisLetsGo/Tetris/tool"
	"math/rand"
	"sync"
	"time"

	//"time"
)

func main() {
	arr := make([][]int,25)
	for i:=0 ; i<25; i++ {
		arr[i] = make([]int,10)
	}

	for i:=0; i<10; i++ {
		arr[0][i] = 8
	}

	fmt.Println("Press ESC to quit")

	for {
		randShape := rand.Intn(7) + 1
		point := []int{24, 5}
		point[0] = 24
		point[1] = 5

		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		defer func() {
			_ = keyboard.Close()
		}()

		wait := sync.WaitGroup{}
		wait.Add(3)
		ch1 := make(chan string)
		ch2 := make(chan string, 1)
		ch3 := make(chan bool)
		ch4 := make(chan bool)
		go eventHandler(ch1, ch2, arr, randShape, 0, point, wait)
		go down(ch1, ch3, wait) // cd Desktop/mspring03/project/TetrisLetsGo/Tetris/
		go keyListener(ch1, ch4, wait)
		fmt.Println(1234)
		<- ch2

		ch3 <- true
		fmt.Println(1234)

		ch4 <- true
		fmt.Println(1234)


		//wait.Done()
		close(ch1)
		close(ch2)
		close(ch3)
		close(ch4)
	}
}

func eventHandler(ch1 chan string, ch2 chan string, arr [][]int, randomShape int, shapeNum int, point []int, wait sync.WaitGroup) {
	for {
		switch input := <-ch1; input {
		case "up":
			shapeNum %= 4
			shape := tool.Shapes(randomShape, shapeNum+1) // {0, 0}, {0, -1}, {-1, 0}, {-1, -1}
			if (arr[point[0]+shape[0][0]][point[1]+shape[0][1]] != 0) {
				break
			}
			if (arr[point[0]+shape[1][0]][point[1]+shape[1][1]] != 0) {
				break
			}
			if (arr[point[0]+shape[2][0]][point[1]+shape[2][1]] != 0) {
				break
			}
			if (arr[point[0]+shape[3][0]][point[1]+shape[3][1]] != 0) {
				break
			}
			arr[point[0]+shape[0][0]][point[1]+shape[0][1]] = randomShape
			arr[point[0]+shape[1][0]][point[1]+shape[1][1]] = randomShape
			arr[point[0]+shape[2][0]][point[1]+shape[2][1]] = randomShape
			arr[point[0]+shape[3][0]][point[1]+shape[3][1]] = randomShape
			tool.Print(arr)
			arr[point[0]+shape[0][0]][point[1]+shape[0][1]] = 0
			arr[point[0]+shape[1][0]][point[1]+shape[1][1]] = 0
			arr[point[0]+shape[2][0]][point[1]+shape[2][1]] = 0
			arr[point[0]+shape[3][0]][point[1]+shape[3][1]] = 0
			shapeNum += 1
			break
		case "left":
		case "right":
		case "down":
			shape := tool.Shapes(randomShape, shapeNum+1)
			if point[0]+shape[0][0] == 0 || point[0]+shape[1][0] == 0 || point[0]+shape[2][0] == 0 || point[0]+shape[3][0] == 0 {
				arr[point[0]+shape[0][0]+1][point[1]+shape[0][1]] = randomShape
				arr[point[0]+shape[1][0]+1][point[1]+shape[1][1]] = randomShape
				arr[point[0]+shape[2][0]+1][point[1]+shape[2][1]] = randomShape
				arr[point[0]+shape[3][0]+1][point[1]+shape[3][1]] = randomShape
				tool.Print(arr)
				ch2 <- "close"
				wait.Done()
				return
			}
			if arr[point[0]+shape[0][0]][point[1]+shape[0][1]] != 0 {
				arr[point[0]+shape[0][0]+1][point[1]+shape[0][1]] = randomShape
				arr[point[0]+shape[1][0]+1][point[1]+shape[1][1]] = randomShape
				arr[point[0]+shape[2][0]+1][point[1]+shape[2][1]] = randomShape
				arr[point[0]+shape[3][0]+1][point[1]+shape[3][1]] = randomShape
				tool.Print(arr)
				ch2 <- "close"
				wait.Done()
				return
			}
			if arr[point[0]+shape[1][0]][point[1]+shape[1][1]] != 0 {
				arr[point[0]+shape[0][0]+1][point[1]+shape[0][1]] = randomShape
				arr[point[0]+shape[1][0]+1][point[1]+shape[1][1]] = randomShape
				arr[point[0]+shape[2][0]+1][point[1]+shape[2][1]] = randomShape
				arr[point[0]+shape[3][0]+1][point[1]+shape[3][1]] = randomShape
				tool.Print(arr)
				ch2 <- "close"
				wait.Done()
				return
			}
			if arr[point[0]+shape[2][0]][point[1]+shape[2][1]] != 0 {
				arr[point[0]+shape[0][0]+1][point[1]+shape[0][1]] = randomShape
				arr[point[0]+shape[1][0]+1][point[1]+shape[1][1]] = randomShape
				arr[point[0]+shape[2][0]+1][point[1]+shape[2][1]] = randomShape
				arr[point[0]+shape[3][0]+1][point[1]+shape[3][1]] = randomShape
				tool.Print(arr)
				ch2 <- "close"
				wait.Done()
				return
			}
			if arr[point[0]+shape[3][0]][point[1]+shape[3][1]] != 0 {
				arr[point[0]+shape[0][0]+1][point[1]+shape[0][1]] = randomShape
				arr[point[0]+shape[1][0]+1][point[1]+shape[1][1]] = randomShape
				arr[point[0]+shape[2][0]+1][point[1]+shape[2][1]] = randomShape
				arr[point[0]+shape[3][0]+1][point[1]+shape[3][1]] = randomShape
				tool.Print(arr)
				ch2 <- "close"
				wait.Done()
				return
			}
			arr[point[0]+shape[0][0]][point[1]+shape[0][1]] = randomShape
			arr[point[0]+shape[1][0]][point[1]+shape[1][1]] = randomShape
			arr[point[0]+shape[2][0]][point[1]+shape[2][1]] = randomShape
			arr[point[0]+shape[3][0]][point[1]+shape[3][1]] = randomShape
			tool.Print(arr)
			arr[point[0]+shape[0][0]][point[1]+shape[0][1]] = 0
			arr[point[0]+shape[1][0]][point[1]+shape[1][1]] = 0
			arr[point[0]+shape[2][0]][point[1]+shape[2][1]] = 0
			arr[point[0]+shape[3][0]][point[1]+shape[3][1]] = 0
			point[0] -= 1
			break
		}
	}
}

func down(ch1 chan string, ch3 chan bool, wait sync.WaitGroup) {
	for {
		select {
		case <- ch3:
			wait.Done()
			return
		default:
			time.Sleep(time.Duration(1) * time.Second)
			ch1 <- "down"
		}
	}
}

func keyListener(ch1 chan string, ch4 chan bool, wait sync.WaitGroup) {
	for {
		select {
		case <- ch4:
			wait.Done()
			return
		default:

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

			}
		}
	}
}