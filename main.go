package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"rest_api_simple/Config"
	"rest_api_simple/Models"
	"rest_api_simple/Routes"
)

var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildCOnfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}