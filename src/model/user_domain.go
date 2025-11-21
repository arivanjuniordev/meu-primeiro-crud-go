package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/arivanjuniordev/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/arivanjuniordev/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func NewUserDomain(email string, password string, name string, age int8) UserDomainInterface {
	return &userDomain{
		email,
		password,
		name,
		age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

func (ud *userDomain) EncriptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))

	logger.Info("User password encrypted",
		zap.String("journey", "createUser encriptPassword"),
		zap.String("password", ud.password),
	)

}
