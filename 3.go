package main

import "fmt"

func main() {
	result := make (chan int,1)
	result <-10
	select {
	case x := <-result:
		fmt.Println("Получено:",x)
	default:
		fmt.Println("Данных нет")
	}
	select {
	case x := <-result:
		fmt.Println("Получено:",x)
	default:
		fmt.Println("Данных нет")
	}
}