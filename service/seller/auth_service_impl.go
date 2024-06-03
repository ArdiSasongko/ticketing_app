package seller

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/repository/seller"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type SellerServiceImpl struct {
	repository   seller.SellerRepository
	tokenUseCase helper.TokenUseCase
}

func NewSellerService(repository seller.SellerRepository, token helper.TokenUseCase) *SellerServiceImpl {
	return &SellerServiceImpl{
		repository:   repository,
		tokenUseCase: token,
	}
}

func (service *SellerServiceImpl) SaveSeller(request SellerServiceRequest) (map[string]interface{}, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if errHash != nil {
		return nil, errHash
	}

	request.Password = string(passHash)
	newseller := domain.Sellers{
		Name:     request.Name,
		Password: request.Password,
		Email:    request.Email,
	}

	saveSeller, errSaveSeller := service.repository.SaveSeller(newseller)
	if errSaveSeller != nil {
		return nil, errSaveSeller
	}

	data := helper.ResponseToJson{
		"name ": saveSeller.Name,
		"email": saveSeller.Email,
	}
	return data, nil
}

func (service *SellerServiceImpl) LoginSeller(email string, password string) (map[string]interface{}, error) {
	seller, err := service.repository.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("Email tidak ditemukan")
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(seller.Password), []byte(password))
	if errPass != nil {
		return nil, errors.New("Password Salah")
	}

	expiredTime := time.Now().Local().Add(1 * time.Hour)

	claims := helper.JwtCustomClaims{
		ID:    strconv.Itoa(seller.SellerID),
		Name:  seller.Name,
		Email: seller.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest-gorm",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	token, errToken := service.tokenUseCase.GenerateAccessToken(claims)
	if errToken != nil {
		return nil, errors.New("ada kesalahan generate token")
	}

	return map[string]interface{}{"token": token}, nil
}
