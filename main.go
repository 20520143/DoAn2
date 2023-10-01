package main

import (
	"projectDemo/controller"
	"projectDemo/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	dsn := "root:@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	controller.ConnectDB(DB)
	routes.AuthRoutes(e)
	routes.StudentClassRoutes(e)
	routes.StudentRoutes(e)
	routes.ClassRoutes(e)
	e.Logger.Fatal(e.Start(":9090"))
}
