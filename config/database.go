package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"playgofiber/models"
)

func InitDB() *gorm.DB {
	dsn := concateConfigureDb();
	
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	PanicHandler(err)

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Book{})

	return db
}

func concateConfigureDb() string {
	concateSring := GetConfigEnv("MYSQL_USERNAME") + ":" +
					GetConfigEnv("MYSQL_PASSWORD") + "@tcp(" +
					GetConfigEnv("MYSQL_HOST") + ":" +
					GetConfigEnv("MYSQL_PORT") + ")/" +
					GetConfigEnv("MYSQL_DB") + "?charset=utf8mb4&parseTime=True&loc=Local"

	return concateSring
}



