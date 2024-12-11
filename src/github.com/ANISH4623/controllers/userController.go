package controllers

import (
	"github.com/ANISH4623/database"
	"github.com/ANISH4623/models"
	"github.com/gin-gonic/gin"
	validate "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func UserResponse(user models.User) models.User {
	return models.User{
		Model:     gorm.Model{},
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		Products:  nil,
	}

}
func CreateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var existingUser models.User
		var user models.User
		if err := context.Bind(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validate.New().Struct(user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		database.Database.Db.Find(&existingUser, "username = ?", user.Username)
		if existingUser.Username != "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
		database.Database.Db.Find(&existingUser, "email = ?", user.Email)
		if existingUser.ID != 0 {
			context.JSON(http.StatusBadRequest, gin.H{"Error": "User Already Exists "})
			return
		}
		password, _ := hashPassword(user.Password)
		user.Password = password
		database.Database.Db.Create(&user)
		var ResponseUser = UserResponse(user)
		context.JSON(http.StatusCreated, ResponseUser)
	}
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
