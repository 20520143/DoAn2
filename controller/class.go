package controller

import (
	"net/http"
	"projectDemo/model"
	"projectDemo/service"

	"github.com/labstack/echo/v4"
)

type TeacherStudent struct {
	Teacher  model.Teacher
	Students []model.Student
}

var class = service.NewClass()

func GetTecherStudent(c echo.Context) error {
	data := make(map[string]interface{})
	classID := c.QueryParam("ClassID")
	item, err := class.GetTeacherStudent(classID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, FailRes(400, err.Error()))
	}
	data[classID] = item
	return c.JSON(http.StatusOK, SuccessRes(data))
}

func GetTeacherStudentAll(c echo.Context) error {
	data := make(map[string]interface{})
	item, err := class.GetTeacherStudentAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, FailRes(400, err.Error()))
	}
	data["list"] = item
	return c.JSON(http.StatusOK, SuccessRes(data))
}

func Create(c echo.Context) error {
	data := make(map[string]interface{})
	var req model.Class
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, FailRes(400, err.Error()))
	}
	err = class.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway, FailRes(502, err.Error()))
	}
	data["state"] = "success"
	return c.JSON(http.StatusOK, SuccessRes(data))
}
