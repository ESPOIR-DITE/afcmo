package user

import "time"

type Users struct {
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}
type UserImage struct {
	Id      string    `json:"id"`
	ImageId string    `json:"imageId"`
	Date    time.Time `json:"date"`
}
type UserGender struct {
	Id       string `json:"id"`
	GenderId string `json:"genderId"`
	Email    string `json:"email"`
}
type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Stat     string `json:"stat"`
	Date     string `json:"date"`
}
type Gender struct {
	Id     string `json:"id"`
	Gender string `json:"gender"`
}
