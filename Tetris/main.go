package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/mspring03/TetrisLetsGo/Tetris/tool"
	"math/rand"
	"os"
	"time"

	//"time"
)

func main() {
	arr := make([][]int,30)
	for i:=0 ; i<30; i++ {
		arr[i] = make([]int,10)
	}

	for i:=0; i<10; i++ {
		arr[0][i] = 8
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	point := []int{24, 5}
	for {
		randShape := rand.Intn(7) + 1
		b := false
		shapeNum := 0
		point[0] = 24
		point[1] = 5

		go func() {
			for {
				time.Sleep(time.Duration(1) * time.Second)
				b = eventHandler("down", arr, randShape, &shapeNum, point)
				if b {
					return
				}
			}
		}()

		for {
			_, key, _ := keyboard.GetKey()
			if b {
				break
			}
			switch key {
			case keyboard.KeyArrowRight:
				_ = eventHandler("right", arr, randShape, &shapeNum, point)
			case keyboard.KeyArrowUp:
				_ = eventHandler("up", arr, randShape, &shapeNum, point)
			case keyboard.KeyArrowLeft:
				_ = eventHandler("left", arr, randShape, &shapeNum, point)
			case keyboard.KeyArrowDown:
				b = eventHandler("down", arr, randShape, &shapeNum, point)
			case keyboard.KeyEsc:
				os.Exit(3)
			}
		}
	}
}

func eventHandler(c string, arr [][]int, randomShape int, shapeNum *int, point []int) bool {
	carr := arr
	switch c {
	case "up":
		*shapeNum = (*shapeNum + 1)%4
		shape := tool.Shapes(randomShape, *shapeNum+1) // {0, 0}, {0, -1}, {-1, 0}, {-1, -1}
		if point[0]+shape[0][0] == 0 || point[0]+shape[1][0] == 0 || point[0]+shape[2][0] == 0 || point[0]+shape[3][0] == 0 {
			return true
		}
		if arr[point[0]+shape[0][0]][point[1]+shape[0][1]] != 0 {
			return false
		}
		if arr[point[0]+shape[1][0]][point[1]+shape[1][1]] != 0 {
			return false
		}
		if arr[point[0]+shape[2][0]][point[1]+shape[2][1]] != 0 {
			return false
		}
		if arr[point[0]+shape[3][0]][point[1]+shape[3][1]] != 0 {
			return false
		}
		carr[point[0]+shape[0][0]][point[1]+shape[0][1]] = randomShape
		carr[point[0]+shape[1][0]][point[1]+shape[1][1]] = randomShape
		carr[point[0]+shape[2][0]][point[1]+shape[2][1]] = randomShape
		carr[point[0]+shape[3][0]][point[1]+shape[3][1]] = randomShape
		tool.Print(arr)
		carr[point[0]+shape[0][0]][point[1]+shape[0][1]] = 0
		carr[point[0]+shape[1][0]][point[1]+shape[1][1]] = 0
		carr[point[0]+shape[2][0]][point[1]+shape[2][1]] = 0
		carr[point[0]+shape[3][0]][point[1]+shape[3][1]] = 0
		return false
	case "left":
	case "right":
	case "down":
		shape := tool.Shapes(randomShape, *shapeNum+1)
		if point[0]+shape[0][0] == 0 || point[0]+shape[1][0] == 0 || point[0]+shape[2][0] == 0 || point[0]+shape[3][0] == 0 {
			arr[point[0]+shape[0][0]+1][point[1]+shape[0][1]] = randomShape
			arr[point[0]+shape[1][0]+1][point[1]+shape[1][1]] = randomShape
			arr[point[0]+shape[2][0]+1][point[1]+shape[2][1]] = randomShape
			arr[point[0]+shape[3][0]+1][point[1]+shape[3][1]] = randomShape
			tool.Print(arr)
			return true
		}
		if arr[point[0]+shape[0][0]][point[1]+shape[0][1]] != 0 {
			arr[point[0]+shape[0][0]+1][point[1]+shape[0][1]] = randomShape
			arr[point[0]+shape[1][0]+1][point[1]+shape[1][1]] = randomShape
			arr[point[0]+shape[2][0]+1][point[1]+shape[2][1]] = randomShape
			arr[point[0]+shape[3][0]+1][point[1]+shape[3][1]] = randomShape
			tool.Print(arr)
			return true
		}
		if arr[point[0]+shape[1][0]][point[1]+shape[1][1]] != 0 {
			arr[point[0]+shape[0][0]+1][point[1]+shape[0][1]] = randomShape
			arr[point[0]+shape[1][0]+1][point[1]+shape[1][1]] = randomShape
			arr[point[0]+shape[2][0]+1][point[1]+shape[2][1]] = randomShape
			arr[point[0]+shape[3][0]+1][point[1]+shape[3][1]] = randomShape
			tool.Print(arr)
			return true
		}
		if arr[point[0]+shape[2][0]][point[1]+shape[2][1]] != 0 {
			arr[point[0]+shape[0][0]+1][point[1]+shape[0][1]] = randomShape
			arr[point[0]+shape[1][0]+1][point[1]+shape[1][1]] = randomShape
			arr[point[0]+shape[2][0]+1][point[1]+shape[2][1]] = randomShape
			arr[point[0]+shape[3][0]+1][point[1]+shape[3][1]] = randomShape
			tool.Print(arr)
			return true
		}
		if arr[point[0]+shape[3][0]][point[1]+shape[3][1]] != 0 {
			arr[point[0]+shape[0][0]+1][point[1]+shape[0][1]] = randomShape
			arr[point[0]+shape[1][0]+1][point[1]+shape[1][1]] = randomShape
			arr[point[0]+shape[2][0]+1][point[1]+shape[2][1]] = randomShape
			arr[point[0]+shape[3][0]+1][point[1]+shape[3][1]] = randomShape
			tool.Print(arr)
			return true
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
		return false
	}
	return false
}