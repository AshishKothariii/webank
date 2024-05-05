package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AshishKothariii/webank/pkg/models"
	"github.com/AshishKothariii/webank/pkg/repository"
	"github.com/AshishKothariii/webank/pkg/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
type UserController struct{
	session *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController{
return &UserController{s}
}
func (uc UserController) CreateAccount(w http.ResponseWriter, r *http.Request) {
        var request request.CreateAccountRequest
        err := json.NewDecoder(r.Body).Decode(&request)
        if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
        }
       acc,_:=models.NewAccount(request)
       ans := repository.CreateAccount(acc)
       fmt.Fprintf(w,"{message : %s}", ans)

}
func (uc UserController) LoginAccount(w http.ResponseWriter, r *http.Request) {
        var request request.LoginRequest
        err :=json.NewDecoder(r.Body).Decode(&request)
        if err != nil{
                http.Error(w,err.Error(),http.StatusBadRequest)
                return
        }
     existingAccount := models.Account{}
        filter := bson.M{"email": request.Email}
        err = uc.session.Database("test").Collection("accounts").FindOne(context.TODO(), filter).Decode(&existingAccount)
        if err != nil {
                // An account with the same email already exists
                http.Error(w, "Email doesn't exists", http.StatusConflict)
                return
        }
       ans :=models.CheckPass(&existingAccount,request.Password)
     fmt.Println(ans)
     if ans==false{
        return
     }
  
      cookie := &http.Cookie{
        Name:  "auth_token",
        Value: fmt.Sprintf("%s:%s", existingAccount.Name, existingAccount.Email),
        Path:  "/",
        MaxAge: 60*60*24*1, // 30 days
        HttpOnly: false,
        Secure: true,
    }
    http.SetCookie(w, cookie)
}

func (uc UserController) Getbalance(w http.ResponseWriter,r *http.Request){
cookie,err :=r.Cookie("auth_token")
if err!=nil{
        fmt.Println("token not authenticated")
        return
}
  parts := strings.Split(cookie.Value, ":")
    if len(parts) != 2 {
        // Handle the case where the cookie value is malformed
    }
    _, email := parts[0], parts[1]
     existingAccount :=models.Account{}
    filter := bson.M{"email": email}
    err = uc.session.Database("test").Collection("accounts").FindOne(context.TODO(), filter).Decode(&existingAccount)
    json.NewEncoder(w).Encode(existingAccount.Balance)
}

