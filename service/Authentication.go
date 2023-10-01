package service

import (
	"fmt"
	"net/http"
	"time"

	"projectDemo/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(db *gorm.DB) {
	DB = db
}

type authen struct{}

func NewAuthen() IAuthen {
	return &authen{}
}

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}
		secret := []byte("my-secret-key")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Kiểm tra phương thức ký của token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secret), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id := claims["id"].(string)
			name := claims["name"].(string)
			expiration := claims["exp"].(float64)

			// Kiểm tra thời gian hết hạn
			if int64(expiration) < time.Now().Unix() {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token has expired")
			}
			var user model.Users
			result := DB.Where("id=? AND name=?", id, name).First(&user)
			if result.RowsAffected == 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token does not exist")
			}
			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}
	}
}

func createToken(id string, name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Tạo thông tin trong Payload (Claim)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Thời gian hết hạn: 72 giờ sau

	// Tạo chuỗi bí mật (secret key)
	secretKey := []byte("my-secret-key")

	// Tạo chữ ký bằng cách ký token bằng secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (authen *authen) UserAuthentication(req model.Users) (string, error) {
	var user model.Users
	result := DB.Where("email=? AND password=?", req.Email, req.Password).First(&user)
	if result.RowsAffected == 0 {
		return "", nil
	} else {
		t, err := createToken(user.Id, user.Name)
		if err != nil {
			return "", err
		}
		return t, nil
	}
}
func (authen *authen) CreateUser(req model.Users) (string, error) {
	id := uuid.New().String()
	req.Id = id
	result := DB.Create(&req)
	if result.Error != nil {
		return "", result.Error
	} else {
		t, err := createToken(req.Id, req.Name)
		if err != nil {
			return "", err
		}
		return t, nil
	}

}
