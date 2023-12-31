package controllers

import (
	"context"
	"golang-jwt-project/database"
	"golang-jwt-project/helpers"
	"golang-jwt-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func HashPassword() {

}

func VerifyPassword() {

}

func GetUsers() gin.HandlerFunc {
}

// * This GetUser() is used to fetch UserInformation from db after validating the User Details
// * return User after Success Else Return Error
func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")

		if err := helpers.MatchUserTypeToUid(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		//* Find the User in the Mongo Collection
		var user models.User
		err := userCollection.FindOne(c, bson.M{"user_id": userId}).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func Login() gin.HandlerFunc {

}

func Signup() {

}
