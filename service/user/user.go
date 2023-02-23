package user

import (
	"errors"
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"github.com/adeben33/vehicleParkingApi/pkg/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func SignUpUser(user model.User) (model.UserRes, string, error) {
	//check if the user already exists this is checked by email which is unique based of validator
	_, err := mongodb.FindUser(user.Email)
	if err == nil {
		return model.UserRes{}, fmt.Sprintf("User Already exists"), errors.New("user already exist in database")
	}
	//Harsh the password
	harsh, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(harsh)
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().Local().Format(time.DateTime)
	user.CreatedAt = time.Now().Local().Format(time.DateTime)
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
	}
	return response, fmt.Sprintf("User saved successfully"), nil
}
