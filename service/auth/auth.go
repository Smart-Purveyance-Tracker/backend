package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID uint64) (string, error)
	UserIDFromToken(token string) (uint64, error)
}

type JWTService struct {
	accessSecret []byte
}

func NewJWTService(accessSecret []byte) *JWTService {
	return &JWTService{
		accessSecret: accessSecret,
	}
}

func (s *JWTService) GenerateToken(userID uint64) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString(s.accessSecret)
}

func (s *JWTService) UserIDFromToken(token string) (uint64, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.accessSecret, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("no map claims")
	}

	userID, ok := claims["user_id"].(uint64)
	if !ok {
		return 0, errors.New("not valid user id type")
	}

	return userID, nil
}
