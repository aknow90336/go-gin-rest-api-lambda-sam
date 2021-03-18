package src

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"os"
)

var DBHelper *gorm.DB
var err error

func init() {
	connectionString, _:= os.LookupEnv("Prod_Connection_String")

	DBHelper, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	sqlDB, err := DBHelper.DB()
	
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(100)
}
