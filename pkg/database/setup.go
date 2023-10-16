package database

import (
	"fmt"
	"os"

	"github.com/dawamr/resume-cv-2023/app/models"
	"github.com/dawamr/resume-cv-2023/pkg/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() {
	var err any
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		env.GetEnv("DB_USER", os.Getenv("DB_USER")),
		env.GetEnv("DB_PASSWORD", os.Getenv("DB_PASSWORD")),
		env.GetEnv("DB_HOST", os.Getenv("DB_HOST")),
		env.GetEnv("DB_PORT", os.Getenv("DB_PORT")),
		env.GetEnv("DB_NAME", os.Getenv("DB_NAME")),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.User{})
}
