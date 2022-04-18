package jwt_service

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
)

type DecodedToken struct {
	Id     int    `json:"id"`
	Iat    int    `json:"iat"`
	Role   string `json:"role"`
	UserId string `json:"userId"`
	Email  string `json:"email"`
	Iss    string `json:"iss"`
}

const Authorization = "Authorization"

func GenerateToken(claims *jwt.Token, secret string) string {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ := claims.SignedString(hmacSecret)

	return token
}

func VerifyToken(token string, secret string) *DecodedToken {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil
	}

	if !decoded.Valid {
		return nil
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	json.Unmarshal(jsonString, &decodedToken)

	return &decodedToken
}
