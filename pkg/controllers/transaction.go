package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AshishKothariii/webank/pkg/models"
	"github.com/AshishKothariii/webank/pkg/request"
	"go.mongodb.org/mongo-driver/bson"
)

func (uc UserController) SendMoney(w http.ResponseWriter,r *http.Request){
 var request request.SendRequest
        err := json.NewDecoder(r.Body).Decode(&request)
	if err!=nil{
		return
	}
	cookie,err :=r.Cookie("auth_token")
if err!=nil{
        fmt.Println("token not authenticated")
        return
}
  parts := strings.Split(cookie.Value, ":")
   
    _, email := parts[0], parts[1]
	fmt.Print(email)
	var sender_acc models.Account
		uc.session.Database("test").Collection("accounts").FindOne(context.TODO(),bson.M{"email": email}).Decode(&sender_acc)
    var to_acc models.Account
		uc.session.Database("test").Collection("accounts").FindOne(context.TODO(),bson.M{"email": request.ToAccount}).Decode(&to_acc)

	
	   uc.session.Database("test").Collection("accounts").FindOneAndUpdate(context.TODO(),bson.M{"email": sender_acc.Email},bson.M{"$set": bson.M{"balance":sender_acc.Balance-request.Amount}})
      fmt.Println("sent from sender account email: ",request.Amount,sender_acc.Email)
	 uc.session.Database("test").Collection("accounts").FindOneAndUpdate(context.TODO(),bson.M{"email": to_acc.Email},bson.M{"$set": bson.M{"balance":to_acc.Balance+request.Amount}})
      fmt.Println("received from sender account email: ",request.Amount,to_acc.Email)

}
func (uc UserController) GetTransactions(w http.ResponseWriter,r *http.Request){
	
}