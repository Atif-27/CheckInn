package types

import (
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	BCRYPT_COST     = 12
	minLenFirstname = 2
	minLenLastname  = 2
	minLenPassword  = 7
)

type UpdateUserParams struct {
	Firstname string `bson:"firstName" json:"firstName"`
	Lastname  string `bson:"lastName" json:"lastName"`
}
type CreateUserParams struct {
	Firstname string `bson:"firstName" json:"firstName"`
	Lastname  string `bson:"lastName" json:"lastName"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password,omitempty"`
}

func (params CreateUserParams) Validate() []string {
	log.Info("enter Validate()")
	errors := []string{}

	if len(params.Firstname) < minLenFirstname {
		errors = append(errors, fmt.Sprintf("firstName should be atleast %d characters", minLenFirstname))
	}
	if len(params.Lastname) < minLenLastname {
		errors = append(errors, fmt.Sprintf("lastName should be atleast %d characters", minLenLastname))
	}
	if len(params.Password) < minLenPassword {
		errors = append(errors, fmt.Sprintf("password should be atleast %d characters", minLenPassword))
	}
	if !isEmailValid(params.Email) {
		errors = append(errors, fmt.Sprintf("email %s is invalid", params.Email))
	}
	log.Info("exit Validate()")
	log.Info("errors = ", errors)
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Firstname string             `bson:"firstName" json:"firstName"`
	Lastname  string             `bson:"lastName" json:"lastName"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), BCRYPT_COST)
	if err != nil {
		return nil, err
	}
	return &User{
		Firstname: params.Firstname,
		Lastname:  params.Lastname,
		Email:     params.Email,
		Password:  string(encpw),
	}, nil
}
