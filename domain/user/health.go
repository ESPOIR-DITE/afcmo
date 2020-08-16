package user

import "time"

type Health struct {
	Id          string `json:"id"`
	Health      string `json:"health"`
	Description string `json:"description"`
}
type UserHealth struct {
	Id          string    `json:"id"`
	Email       string    `json:"email"`
	HealthId    string    `json:"healthId"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
