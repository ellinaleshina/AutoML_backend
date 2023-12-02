package service

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateJwt(login string) string {
	signingKey := []byte("key")
	user := ur.GetByEmail(login)
	// Create claims with multiple fields populated
	claims := Claims{
		login,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        strconv.Itoa(user.Id),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		return ""
	}
	return ss
}
