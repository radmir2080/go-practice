package repository

import (
	"fmt"
	"log"

	"smart-itern/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := "host=localhost user=postgres password=8700836 dbname=smart_intern port=5555 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных :", err)
	}

	fmt.Println("Успешно подключено к базе данных PostgreSQL")

	err = db.AutoMigrate(
		&domain.User{},
		&domain.Student{},
		&domain.Company{},
		&domain.Internship{},
	)
	if err != nil {
		log.Fatal("Ошибка автомиграции: ", err)
	}

	fmt.Println("База данных успешно синхронизирована!")

	return db
}
