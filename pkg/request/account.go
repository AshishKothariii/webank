package request

type CreateAccountRequest struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Gender    string `json:"gender" bson:"gender"`
	Password  string `json:"password" bson:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type LoginResponse struct {
	Email string `json:"email" bson:"email"`
	Token string `json:"token" bson:"token"`
}
type LogoutRequest struct {
	Email    string `json:"email"`
	IsLogout bool   `json:"logout"`
}
