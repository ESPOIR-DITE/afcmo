package nutrition

import "time"

type Nutrition struct {
	Id            string  `json:"id"`
	NutritionType string  `json:"nutritionType"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
}
type NutritionImage struct {
	Id          string `json:"id"`
	ImageId     string `json:"imageId"`
	Description string `json:"description"`
}
type NutritionType struct {
	Id            string    `json:"id"`
	NutritionType string    `json:"nutritionType"`
	Date          time.Time `json:"date"`
}
type NutritionUserPlan struct {
	Id             string    `json:"id"`
	Email          string    `json:"email"`
	NutritionId    string    `json:"nutrition_id"`
	SubscriptionId string    `json:"subscriptionId"`
	Date           time.Time `json:"date"`
}
type NutritionVideo struct {
	Id          string `json:"id"`
	VideoId     string `json:"videoId"`
	NutritionId string `json:"nutritionId"`
	Description string `json:"description"`
}
