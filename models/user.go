package models

import (
	// "rest-api-gin-gonic-mongo/db"
	"rest-api-gin-gonic-mongo/forms"
	"time"
	"log"
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email              string `idx:"{email},unique" json:"email" binding:"required"`
	Password           string `json:"password" binding:"required"`
	Name               string `json:"name"`
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
	VerifiedAt *time.Time
}


type UserModel struct{}

// var server = "127.0.0.1"

// var dbConnect = db.NewConnection(server)

func (m *UserModel) Register(data forms.RegisterUserCommand) error {
	collection := dbConnect.Use("test-mgo", "user")
	err := collection.Insert(bson.M{"email": data.Email, "password": data.Password, "name": data.Name})
	return err
}


func (m *UserModel) GetEmail(email string) (user User, err error) {
	collection := dbConnect.Use("test-mgo", "user")
	err = collection.Find(bson.M{"email": email}).One(&user)
	return user, err
}




func (m *UserModel) GetJwtToken(email string) (tokenString string,err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": string(email),
	})
	log.Println(token)

	secretKey := "qwertyy"
	tokenString, err = token.SignedString([]byte(secretKey))
	return tokenString, err
}


