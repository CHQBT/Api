package auth

import (
	"milliy/config"
	"milliy/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	tokenMain = config.Load().Token.ACCES_KEY
)

func GeneratedRefreshJWTToken(req *model.User) (string, error) {
	token := *jwt.New(jwt.SigningMethodHS256)
	//payload
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = req.ID
	claims["role"] = req.Role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().AddDate(0, 6, 0).Unix()

	newToken, err := token.SignedString([]byte(tokenMain))
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func ValidateRefreshToken(tokenStr string) (bool, error) {
	_, err := ExtractRefreshClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractRefreshClaim(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenMain), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return &claims, nil
}

func GetUserInfoFromRefreshToken(refreshTokenString string) (string, string, error) {
	claims, err := ExtractRefreshClaim(refreshTokenString)
	if err != nil {
		return "", "", err
	}

	userID := (*claims)["user_id"].(string)
	Role := (*claims)["role"].(string)

	return userID, Role, nil
}
