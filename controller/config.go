package controller

import (
	"projectDemo/service"

	"gorm.io/gorm"
)

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func ConnectDB(DB *gorm.DB) {
	service.ConnectDB(DB)
}

func SuccessRes(data map[string]interface{}) Response {
	return Response{200, "OK", data}
}
func FailRes(status int, err string) Response {
	return Response{status, err, map[string]interface{}{}}
}
