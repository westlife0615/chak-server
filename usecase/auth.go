package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/westlife0615/chak-server/repository"
	"time"
)

type authCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type AuthUsecase interface {
	Login(string) (int, error)
	GenToken(string) (string, error)
}

type authUsecase struct {
	userRepository repository.UserRepository
}

func (authUsecase *authUsecase) Login(email string) (int, error) {

	// findOne
	user, err := authUsecase.userRepository.Get(email)
	if err != nil {
		return 0, err
	}

	if user.Id > -1 {
		return 1, nil
	}

	return 0, nil
	//
}

func (authUsecase *authUsecase) GenToken(email string) (string, error) {
	claims := &authCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "westlife0615",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString([]byte("westlife0615"))
	if err != nil {
		return "", err
	}
	return stringToken, nil
}

func NewAuthUsecase(tr repository.UserRepository) AuthUsecase {
	return &authUsecase{tr}
}
