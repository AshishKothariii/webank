package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AshishKothariii/webank/pkg/router"
	"github.com/joho/godotenv"
)

func main() {
	    err :=godotenv.Load(".env")
     if err!=nil{
		log.Fatal("panic")
	 }
		r :=router.Router()
		http.ListenAndServe(os.Getenv("SERVER_ADDRESS"),r)

}