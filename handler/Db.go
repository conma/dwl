package handler

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)


const url = "postgres://postgres:postgres@localhost:5432/postgres"

func Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db, err
}