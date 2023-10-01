package routes

import (
	"projectDemo/controller"
	"projectDemo/service"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	authRoutes := e.Group("/users")
	authRoutes.POST("/sign_in", controller.SignIn)
	authRoutes.POST("/sign_up", controller.SignUp)
}

func StudentClassRoutes(e *echo.Echo) {
	StudentClassRoutes := e.Group("/class")
	StudentClassRoutes.Use(service.CheckToken)
	StudentClassRoutes.POST("/register", controller.RegisterClass)
	StudentClassRoutes.POST("/delete", controller.DeleteClass)
}

func StudentRoutes(e *echo.Echo) {
	studentRoutes := e.Group("/student")
	studentRoutes.Use(service.CheckToken)
	studentRoutes.GET("/get_class", controller.GetClass)
}

func ClassRoutes(e *echo.Echo) {
	classRoutes := e.Group("/class")
	classRoutes.Use(service.CheckToken)
	classRoutes.GET("/get_teacher_student", controller.GetTecherStudent)
	classRoutes.GET("/get_teacher_student_all", controller.GetTeacherStudentAll)
	classRoutes.POST("/create", controller.Create)
}
