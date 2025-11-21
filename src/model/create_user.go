package model

import (
	"github.com/arivanjuniordev/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/arivanjuniordev/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser model",
		zap.String("journey", "createUser"),
	)

	ud.EncriptPassword()
	return nil
}
