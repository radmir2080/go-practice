package main

import (
	"fmt"
	"smart-itern/internal/repository"
)

func main() {
	fmt.Println("Запуск приложения Smart Itern...")
	db := repository.InitDB()
	if db != nil {
		fmt.Println("Приложение успешно запустилось и готово к работе")
	}
}
