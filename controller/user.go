package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang_app/database"
	"github.com/golang_app/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func New() *UserRepo {
	db := database.InitDB()
	db.AutoMigrate(&model.User{})
	return &UserRepo{DB: db}
}

func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := model.CreateUser(repository.DB, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &user)

}

func (repository *UserRepo) GetUsers(c *gin.Context) {
	var user []model.User
	err := model.GetUsers(repository.DB, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &user)

}

func (repository *UserRepo) GetUserById(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var user model.User
	err := model.GetUserById(repository.DB, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (repository *UserRepo) UpdateUser(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var user model.User
	err := model.GetUserById(repository.DB, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, user)

	err = model.UpdateUser(repository.DB, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepo) DeleteUser(c *gin.Context) {
	var user model.User
	id, _ := c.Params.Get("id")

	err := model.DeleteUser(repository.DB, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
