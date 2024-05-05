package models

import (
	"time"

	request "github.com/AshishKothariii/webank/pkg/request"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Gender   string `json:"gender" bson:"gender"`
	Balance float64 `json:"balance" bson:"balance"`
	EncryptedPassword string `json:"password" bson:"password"`
	CreatedAt time.Time `json:"createdat" bson:"_createdat"`
}
func (a *Account) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

func NewAccount(acc request.CreateAccountRequest) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		Name: acc.FirstName+" "+acc.LastName,
     	Email: acc.Email,
		Gender: acc.Gender,
		Balance: 0.0,
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().Local(),
	}, nil
}
func CheckPass(a *Account,pw string) bool{
	return a.ValidPassword(pw)
}