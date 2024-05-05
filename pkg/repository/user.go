package repository

import (
	"context"
	"os"

	"github.com/AshishKothariii/webank/pkg/models"
	"github.com/AshishKothariii/webank/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateAccount(acc *models.Account) string{
user :=storage.GetSession().Database(os.Getenv("Db_name")).Collection("accounts")
var existingAccount models.Account
user.FindOne(context.Background(),bson.M{
	"email":acc.Email}).Decode(&existingAccount)
if existingAccount.Email==acc.Email{
return "Account exists"
}
user.InsertOne(context.Background(),acc)
return "Account Created"
}