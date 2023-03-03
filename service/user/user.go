package user

import (
	"errors"
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/config"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"github.com/adeben33/vehicleParkingApi/pkg/repository/mongodb"
	"github.com/adeben33/vehicleParkingApi/utility"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func SignUpUser(user model.User) (model.UserRes, string, error) {
	//check if the user already exists this is checked by email which is unique based of validator
	_, count, err := mongodb.FindUser(user.Email)
	if err == nil {
		return model.UserRes{}, fmt.Sprintf("User Already exists "), errors.New("user already exist in database 1")
	}
	if count > 1 {
		return model.UserRes{}, fmt.Sprintf("User Already exists "), errors.New("user already exist in database 1")
	}

	//Harsh the password
	harsh, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(harsh)
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().Local().Format(time.DateTime)
	user.UpdatedAt = time.Now().Local().Format(time.DateTime)
	user.LastLogin = time.Now().Local().Format(time.DateTime)
	user.UserId = user.ID.Hex()

	//Save to the db
	_, err = mongodb.SaveUser(user)
	if err != nil {
		return model.UserRes{}, fmt.Sprintf("Unable to save to the database"), err
	}

	response := model.UserRes{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		VehiclesId: user.VehiclesId,
		LastLogin:  user.LastLogin,
	}
	return response, fmt.Sprintf("User saved successfully"), nil
}

func LoginUser(user model.UserLogin) (model.UserRes, string, string, error) {
	//Get the user from the db
	result, _, _ := mongodb.FindUser(user.Email)

	//	Validate the password
	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {

		return model.UserRes{}, "invalid password", " ", errors.New("invalid password")
	}
	//	get a token
	secretKey := config.GetConfig().Server.Secret
	token := utility.GenerateToken(result, secretKey)
	if err != nil {
		return model.UserRes{}, fmt.Sprintf("unable to create token: %v", err.Error()), " ", err
	}
	//	Update the last logged from the database
	mongodb.SaveUserLastUpdate(user.Email, time.Now())
	//	response this is the user details

	response := model.UserRes{
		FirstName:  result.FirstName,
		LastName:   result.LastName,
		Email:      result.Email,
		VehiclesId: result.VehiclesId,
		LastLogin:  time.Now().String(),
	}
	return response, fmt.Sprintf("user logged in and token generated"), token, nil
}
