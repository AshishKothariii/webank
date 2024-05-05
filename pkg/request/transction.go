package request

type SendRequest struct {
	ToAccount string  `json:"toAccount"`
	Amount    float64 `json:"amount"`
}
type Transaction struct {
}