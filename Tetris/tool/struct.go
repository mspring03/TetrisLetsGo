package tool

func Shapes(shape int, num int) (arr [][]int) {
	switch shape {
	case 1:
		switch num {
		case 1, 2, 3:
		case 4:
			arr := [][]int {
				{0, 0}, {0, -1}, {-1, 0}, {-1, -1},
			}
			return arr
		}
	case 2:
		switch num {
		case 1:
		case 3:
			arr := [][]int {
				{0, -2}, {0, -1}, {0, 0}, {0, 1},
			}
			return arr
		case 2:
		case 4:
			arr := [][]int {
				{1, 0}, {0, 0}, {-1, 0}, {-2, 0},
			}
			return arr
		}
	case 3:
		switch num {
		case 1:
		case 3:
			arr := [][]int {
				{-1, -1}, {-1, 0}, {0, 0}, {0, 1},
			}
			return arr
		case 2:
		case 4:
			arr := [][]int {
				{1, 0}, {0, 0}, {0, 1}, {-1, 1},
			}
			return arr
		}
	case 4:
		switch num {
		case 1:
		case 3:
			arr := [][]int {
				{0, -1}, {0, 0}, {-1, 0}, {-1, 1},
			}
			return arr
		case 2:
		case 4:
			arr := [][]int {
				{1, 1}, {0, 1}, {0, 0}, {-1, 0},
			}
			return arr
		}
	case 5:
		switch num {
		case 1:
			arr := [][]int {
				{-1, -1}, {0, -1}, {0, 0}, {0, 1},
			}
			return arr
		case 2:
			arr := [][]int {
				{1, 0}, {0, 0}, {-1, 0}, {-1, 1},
			}
			return arr
		case 3:
			arr := [][]int {
				{0, -1}, {0, 0}, {0, 1}, {1, 1},
			}
			return arr
		case 4:
			arr := [][]int {
				{1, -1}, {1, 0}, {0, 0}, {-1, 0},
			}
			return arr
		}
	case 6:
		switch num {
		case 1:
			arr := [][]int {
				{0, -1}, {0, 0}, {0, 1}, {-1, 1},
			}
			return arr
		case 2:
			arr := [][]int {
				{-1, 0}, {0, 0}, {1, 0}, {1, 1},
			}
			return arr
		case 3:
			arr := [][]int {
				{1, -1}, {0, -1}, {0, 0}, {0, 1},
			}
			return arr
		case 4:
			arr := [][]int {
				{-1, -1}, {-1, 0}, {0, 0}, {1, 0},
			}
			return arr
		}
	case 7:
		switch num {
		case 1:
			arr := [][]int {
				{0, -1}, {0, 0}, {0, 1}, {-1, 0},
			}
			return arr
		case 2:
			arr := [][]int {
				{1, 0}, {0, 0}, {-1, 0}, {0, 1},
			}
			return arr
		case 3:
			arr := [][]int {
				{0, -1}, {0, 0}, {0, 1}, {1, 0},
			}
			return arr
		case 4:
			arr := [][]int {
				{0, -1}, {0, 0}, {1, 0}, {-1, 0},
			}
			return arr
		}
	}
	return
}