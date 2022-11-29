package authentication

import (
	"backend/src/config"
	enumhelper "backend/src/enum-helpers"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId uint) (string, error) {
	grants := jwt.MapClaims{}
	grants["authorized"] = true
	grants["exp"] = time.Now().Add(time.Hour * 12).Unix()
	grants["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, grants)

	return token.SignedString([]byte(config.JWT_SECRET))
}

func extractToken(request *http.Request) string {
	token := request.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func keyOfVerification(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo inesperado: %v", token.Header["alg"])
	}

	return ([]byte(config.JWT_SECRET)), nil
}

func ValidateToken(request *http.Request) error {
	onlyToken := extractToken(request)

	token, err := jwt.Parse(onlyToken, keyOfVerification)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New(enumhelper.InvalidToken)
}

func ExtractUserId(request *http.Request) (uint, error) {
	onlyToken := extractToken(request)

	token, err := jwt.Parse(onlyToken, keyOfVerification)
	if err != nil {
		return 0, err
	}

	if grants, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", grants["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return uint(userId), nil
	}

	return 0, errors.New(enumhelper.InvalidToken)
}
