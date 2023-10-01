package controller

import (
	"net/http"

	"projectDemo/model"
	"projectDemo/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// JWT
type jwtCustomClaims struct {
	email    string
	password string
	jwt.RegisteredClaims
}

var authen = service.NewAuthen()

func SignIn(c echo.Context) error {
	data := make(map[string]interface{})
	var req model.Users
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, FailRes(400, err.Error()))
	}
	token, err := authen.UserAuthentication(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway, FailRes(502, err.Error()))
	}
	if token == "" {
		return c.JSON(http.StatusUnauthorized, FailRes(401, "Incorrect email or password"))
	} else {
		data["token"] = token
		return c.JSON(http.StatusOK, SuccessRes(data))
	}

}

func SignUp(c echo.Context) error {
	data := make(map[string]interface{})
	var req model.Users
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, FailRes(400, err.Error()))
	}
	token, err := authen.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway, FailRes(502, err.Error()))
	}
	data["token"] = token
	return c.JSON(http.StatusOK, SuccessRes(data))
}
