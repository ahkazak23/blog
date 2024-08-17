package helper

//import (
//	"log"
//	"time"
//
//	"blog/model"
//	"github.com/golang-jwt/jwt/v5"
//)
//
//type CustomClaims struct {
//	Email  string `json:"email"`
//	UserId uint   `json:"user_id"`
//
//	jwt.RegisteredClaims
//}
//
//var secret = []byte("secret")
//
//// GenerateToken creates a new JWT token for a given user
//func GenerateToken(user model.User) (string, error) {
//	claims := CustomClaims{
//		Email:  user.Email,
//		UserId: user.ID,
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Minute)),
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	t, err := token.SignedString(secret)
//	if err != nil {
//		log.Println("Error in token signing:", err)
//		return "", err
//	}
//
//	return t, nil
//}
//
//// ValidateToken parses and validates a JWT token
//func ValidateToken(clientToken string) (*CustomClaims, string) {
//	token, err := jwt.ParseWithClaims(clientToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return secret, nil
//	})
//
//	if err != nil {
//		return nil, err.Error()
//	}
//
//	claims, ok := token.Claims.(*CustomClaims)
//	if !ok {
//		return nil, "Invalid token claims"
//	}
//
//	return claims, ""
//}
//
