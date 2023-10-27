package utils

import (
	"fmt"
	"log"
	"notegin/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(userid int) string {
	secretKey := []byte(os.Getenv("SuperSecretKey"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtUserClaims{
		Userid: userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			Subject:   fmt.Sprint(userid),
		},
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return ""
	}
	return tokenString
}

func GetClaimsFromJwtToken(tokenString string) (jwt.MapClaims, error) {
	secretKey := []byte(os.Getenv("SuperSecretKey"))
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	var claims, ok = token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		log.Println(claims["userid"])
	} else {
		log.Println()
	}

	return claims, nil
}
