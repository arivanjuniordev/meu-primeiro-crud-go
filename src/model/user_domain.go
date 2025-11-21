package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/arivanjuniordev/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/arivanjuniordev/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func NewUserDomain(email string, password string, name string, age int8) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

func (ud *UserDomain) EncriptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))

	logger.Info("User password encrypted",
		zap.String("journey", "createUser encriptPassword"),
		zap.String("password", ud.Password),
	)

}
