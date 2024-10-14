package database

import (
	"fmt"
	"github.com/TheAnonymousDarck/go-practica-7/cmd/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Driver struct {
	Driver *gorm.DB
}

func NewDatabaseDriver() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.MYSQL_DATABASE_URL), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al conectar a la base de datos: ", err)
		return nil, err
	}

	// Migración automática de modelos
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("Error al migrar la base de datos: ", err)
		return nil, err
	}

	return db, nil
}
