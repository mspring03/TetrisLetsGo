package tool

func Pop(arr [][]int, line int) {
	_ = append(arr[:line], arr[line+1:]...)
	_ = append(arr, make([]int,10))
	return
}