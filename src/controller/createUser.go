package controller

import (
	"net/http"

	"github.com/arivanjuniordev/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/arivanjuniordev/meu-primeiro-crud-go/src/configuration/valitadion"
	"github.com/arivanjuniordev/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/arivanjuniordev/meu-primeiro-crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest := valitadion.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfuly",
		zap.String("journey", "createUser"),
	)

	c.String(http.StatusOK, "")

}
