package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Начало")
	result := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		result <- ("Ответ от сервера")
	}()
	select {
	case  <-result:
		fmt.Println("Вперед")
	case <-time.After(2 * time.Second):
		fmt.Println("Время ожидания истекло")
	}
	fmt.Println("Конец программы")
}