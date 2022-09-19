package controllers

import (
	"errors"
	"golang-user-service/helpers"
	"golang-user-service/src/account"
	"golang-user-service/src/account/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	accountService account.UserService
}

func CreateUserHandler(r *gin.Engine, userService account.UserService) {
	userHandler := UserController{userService}

	r.POST("/register", userHandler.addUser)
	r.POST("/login", userHandler.signIn)
	r.GET("/users", helpers.MiddlewareJWTAuthorization(), userHandler.getById)
}

func (e *UserController) addUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]helpers.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = helpers.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	newUser, err := e.accountService.Create(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"userId": newUser.ID, "name": newUser.Name, "username": newUser.Username})
}
func (e *UserController) signIn(c *gin.Context) {
	var userdto models.SignIn
	if err := c.ShouldBindJSON(&userdto); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]helpers.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = helpers.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	user, err := e.accountService.ReadByUsername(userdto.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Username tidak ditemukan"})
		return
	}
	if err := user.CheckPassword(userdto.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		c.Abort()
		return
	}
	token := helpers.GenerateToken(user.ID, user.Name, true)
	c.JSON(http.StatusOK, gin.H{"userId": user.ID, "name": user.Name, "jwt": token})
}
func (e *UserController) getById(c *gin.Context) {
	decoded := c.MustGet("decoded").(*helpers.AuthCustomClaims)

	user, err := e.accountService.ReadById(decoded.Id.String())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Username tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userId": user.ID, "name": user.Name, "username": user.Username})

}
