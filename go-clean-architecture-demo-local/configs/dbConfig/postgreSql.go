package dbConfig

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresSql *gorm.DB

func ConnectDBPostgresSQL() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_PORT"), os.Getenv("POSTGRES_DB_USERNAME"), os.Getenv("POSTGRES_DB_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"))
	PostgresSqlDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Printf("gorm can't open postgres cause: %s", err.Error())
	}

	if PostgresSqlDB == nil {
		return
	}

	sqlDB, err := PostgresSqlDB.DB()

	if sqlDB == nil {
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Minute * 5)

	PostgresSql = PostgresSqlDB
}
