package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/oleiade/reflections.v1"
	"time"
)

type User struct {
	Email        string    `valid:"email,required"`
	First_name   string    `valid:"first_name,required"`
	Second_name  string    `valid:"second_name,required"`
	Password     []byte    `valid:"required"`
	Is_professor bool      `valid:"-"`
	Id           int       `valid:"-"`
	Created_at   time.Time `valid:"-"`
}

func CreateUser(user User) (Response resp) {
	hash, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	utilities.CheckError(err)
	reflections.SetField(&user, "Password", hash)
	_, err = govalidator.ValidateStruct(user)
	utilities.CheckError(err)
}
