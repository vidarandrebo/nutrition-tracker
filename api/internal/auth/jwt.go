package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

type JwtService struct {
	key []byte
}

func NewJwtService() *JwtService {
	return &JwtService{
		key: []byte("this is some key that should be changed"),
	}
}

func (js *JwtService) CreateToken(userID int64) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		// https://datatracker.ietf.org/doc/html/rfc7519#section-4.1
		jwt.MapClaims{
			"iss": "nt.vidarboe.com",
			"sub": userID,
		},
	)
	s, err := t.SignedString(js.key)
	return s, err
}

func (js *JwtService) ValidateToken(tokenString string) (int64, error) {
	extractedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return js.key, nil
	})
	if err != nil {
		return 0, errors.New("failed to validate token")
	}
	id, err := extractedToken.Claims.GetSubject()
	return strconv.ParseInt(id, 10, 64)
}
