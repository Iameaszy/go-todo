package AuthService

import (
	"errors"
	"fmt"
	"time"
	"todo/constants"
	"todo/models"
	"todo/modules/Auth/AuthRepository"

	jwt "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func Login(user *models.User, email, password string) (err error, token string) {
	if err := AuthRepository.GetUser(user, email); err != nil {
		return err, ""
	}
	if user.Email == "" {
		return fmt.Errorf(constants.Errormessages.UserDoesNotExists), ""
	}

	if err = user.CheckPassword(password); err != nil {
		return fmt.Errorf(constants.Errormessages.EmailOrPasswordIncorrect), ""
	}
	if err, token = GenerateJWT(*user); err != nil {
		return err, ""
	}
	return nil, token
}

func Signup(user *models.User) (err error, token string) {
	var dbUser models.User
	if err := AuthRepository.GetUser(&dbUser, user.Email); err != nil {
		return err, ""
	}

	if dbUser.Email != "" {
		return fmt.Errorf(constants.Errormessages.UserAlreadyExists), ""
	}

	if err := user.HashPassword(user.Password); err != nil {
		return err, ""
	}

	if err := AuthRepository.CreateUser(user); err != nil {
		return err, ""
	}

	if err, token = GenerateJWT(*user); err != nil {
		return err, ""
	}

	return nil, token
}

func GenerateJWT(user models.User) (err error, tokenString string) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
