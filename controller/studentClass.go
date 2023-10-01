package controller

import (
	"net/http"
	"projectDemo/model"
	"projectDemo/service"

	"github.com/labstack/echo/v4"
)

var studentClass = service.NewStudentClass()

func RegisterClass(c echo.Context) error {
	data := make(map[string]interface{})
	var req model.StudentClassReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, FailRes(400, err.Error()))
	}
	err = studentClass.RegisterClass(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway, FailRes(502, err.Error()))
	}
	data["State"] = "Success"
	return c.JSON(http.StatusOK, SuccessRes(data))
}

func DeleteClass(c echo.Context) error {
	data := make(map[string]interface{})
	var req model.StudentClassReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, FailRes(400, err.Error()))
	}
	err = studentClass.DeleteClass(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway, FailRes(502, err.Error()))
	}
	data["State"] = "Success"
	return c.JSON(http.StatusOK, SuccessRes(data))
}
