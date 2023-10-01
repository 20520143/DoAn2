package controller

import (
	"net/http"
	"projectDemo/service"

	"github.com/labstack/echo/v4"
)

var student = service.NewStudent()

func GetClass(c echo.Context) error {
	data := make(map[string]interface{})
	studentID := c.QueryParam("StudentID")
	classes, err := student.GetClass(studentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, FailRes(400, err.Error()))
	}
	data["Classes"] = classes
	return c.JSON(http.StatusOK, SuccessRes(data))
}
