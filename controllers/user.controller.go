package controllers

import (
	"net/http"

	"github.com/felipdc/arqbusca-api/models"
	"github.com/felipdc/arqbusca-api/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{UserService: userservice}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newUser, err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"user": newUser})
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	var users []*models.User
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	if len(users) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/", uc.CreateUser)
	// userroute.GET("/:id", uc.GetUser)
	userroute.GET("/", uc.GetAll)
}
