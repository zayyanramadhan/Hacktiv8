package db

import (
	"sesi8-latihan/server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectGorm() *gorm.DB {
	dsn := "root:masukdb@tcp(localhost:3306)/hacktiv8"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Person{})

	return db
}
