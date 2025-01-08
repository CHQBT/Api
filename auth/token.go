package auth

import (
	"milliy/config"
	pb "milliy/generated/api"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GeneratedRefreshJWTToken(req *pb.LoginRes) (string, error) {
	conf := config.Load().Token.ACCES_KEY
	token := *jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = req.Id
	claims["role"] = req.Role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().AddDate(0, 6, 0).Unix()

	newToken, err := token.SignedString([]byte(conf))
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
		return []byte(config.Load().Token.ACCES_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return &claims, nil
}

func GetUserIdFromRefreshToken(token string) (string, string, error) {
	refreshToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return []byte(config.Load().Token.ACCES_KEY), nil })
	if err != nil || !refreshToken.Valid {
		return "", "", err
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", err
	}

	return claims["user_id"].(string), claims["role"].(string), nil
}
