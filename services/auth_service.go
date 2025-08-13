package services

import (
	"context"
	"errors"
	"time"

	"github.com/SatyanarayanSangwal/RootBreakers/config"
	"github.com/SatyanarayanSangwal/RootBreakers/models"
	"github.com/SatyanarayanSangwal/RootBreakers/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var userCollection = config.GetCollection("users")

func RegisterUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if email already exists
	var existing models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existing)
	if err == nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	_, err = userCollection.InsertOne(ctx, user)
	return err
}

func LoginUser(email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
