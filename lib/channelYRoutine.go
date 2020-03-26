package lib

import "fmt"

func sumarSlice(array []int, c chan int) {
	var suma int

	for _, v := range array {
		suma += v
	}

	c <- suma
}

func TestChannelYRoutine() {
	cha := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	mitadArr := len(arr) / 2

	go sumarSlice(arr[:mitadArr], cha)
	go sumarSlice(arr[mitadArr:], cha)

	x := <-cha
	y := <-cha

	fmt.Println(x, y, x+y)
}
