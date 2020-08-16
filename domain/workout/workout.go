package workout

import "time"

type WorkOut struct {
	Id            string    `json:"id"`
	WorkoutTypeId string    `json:"workoutTypeId"`
	Description   string    `json:"description"`
	Date          time.Time `json:"date"`
}
type WorkOutImage struct {
	Id        string `json:"id"`
	WorkoutId string `json:"workOutId"`
	ImageId   string `json:"imageId"`
}
type WorkOutPlan struct {
	Id           string    `json:"id"`
	Email        string    `json:"email"`
	Subscription string    `json:"subscription"`
	Date         time.Time `json:"date"`
}
type WorkoutType struct {
	Id          string    `json:"id"`
	WorkoutType string    `json:"workoutType"`
	Date        time.Time `json:"date"`
}
type WorkVideo struct {
	Id          string `json:"id"`
	WorkOutId   string `json:"workOutId"`
	Description string `json:"description"`
}
