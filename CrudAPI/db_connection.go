package main

import (
	"fmt"

	"gorm.io/driver/mysql" //this automatically get imported from the path (C:\Users\91866\go\pkg\mod\gorm.io\driver)
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:admin@tcp(localhost:3306)/wasimdb"
var err error

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("connection failed :(")
	}
	Database.AutoMigrate(&Employee{})
}
