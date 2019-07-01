package server

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (a *App) BuildDatabase() {

	db, err := gorm.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			a.Config.DBHost,
			a.Config.DBPort,
			a.Config.DBUser,
			a.Config.DBName,
			a.Config.DBPassword,
		),
	)

	if err != nil {
		panic(err)
	}
	a.Database = db
}
