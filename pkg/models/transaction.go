package models

import "time"

type Transaction struct {
	From string `json:"from" bson:"from"`
	To string `json:"to" bson:"to"`
	Amount float64 `json:"amount" bson:"amount"`
	CreatedAt time.Time`json:"createdat" bson:"_createdat"`
}