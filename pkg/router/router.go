package router

import (
	"github.com/AshishKothariii/webank/pkg/controllers"
	"github.com/AshishKothariii/webank/pkg/storage"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	client := storage.GetSession()
    uc :=controllers.NewUserController(client)
	router.HandleFunc("/create",uc.CreateAccount).Methods("POST")
	router.HandleFunc("/login",uc.LoginAccount).Methods("POST")
	router.HandleFunc("/balance",uc.Getbalance).Methods("GET")
	router.HandleFunc("/send",uc.SendMoney).Methods("POST")
	return router
}


