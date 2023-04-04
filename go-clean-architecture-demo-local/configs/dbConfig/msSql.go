package dbConfig

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"gorm.io/driver/sqlserver"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// var MsSql *gorm.DB

// func ConnectDBMsSql() {
// 	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", os.Getenv("MS_DB_USERNAME"), os.Getenv("MS_DB_PASSWORD"), os.Getenv("MS_DB_HOST"), os.Getenv("MS_DB_PORT"), os.Getenv("MS_DB_NAME"))
// 	MsSqlDB, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})

// 	if err != nil {
// 		log.Printf("gorm can't open mssql cause: %s", err.Error())
// 	}

// 	if MsSqlDB == nil {
// 		return
// 	}

// 	sqlDB, err := MsSqlDB.DB()

// 	if sqlDB == nil {
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		return
// 	}

// 	// sqlDB.SetMaxIdleConns(10)
// 	// sqlDB.SetMaxOpenConns(100)
// 	// sqlDB.SetConnMaxLifetime(time.Minute * 5)

// 	MsSql = MsSqlDB
// }
