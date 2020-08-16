package user

import "time"

type PayType struct {
	Id          string `json:"id"`
	PayType     string `json:"payType"`
	Description string `json:"description"`
}
type PayDetail struct {
	Id          string `json:"id"`
	PayTypeId   string `json:"payTypeId"`
	AccountId   string `json:"accountId"`
	BankAccount string `json:"bankAccount"`
	Description string `json:"description"`
}
type NutritionSubscription struct {
	Id          string    `json:"id"`
	Amount      float64   `json:"amount"`
	PaymentType string    `json:"paymentType"`
	Date        time.Time `json:"date"`
}
type WorkoutSubscription struct {
	Id          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	PaymentType string    `json:"paymentType"`
}
