package main

import (
	"fmt"

	// "github.com/Pk05999/library_managment_system/routes"
	// routes "github.com/Pk05999/library_managment_system/routing"
	routes "github.com/Pk05999/library_managment_system/routing"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//Database connection parameter
	dbUser := "root"
	dbPass := "Cloud@710"
	dbName := "libraryManagmentSystem"

	//Constructing the data source name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb&parseTime=True&loc=Local", dbUser, dbPass, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Server started!")
	router := gin.Default()

	//Register book routes
	// routes.AllRoutes(router, db)
	// routes.SetUpBookRoutes(router, db)
	routes.SetUpBookRoutes(router, db)
	router.Run(":8088")
}
