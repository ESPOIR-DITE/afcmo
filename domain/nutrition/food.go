package nutrition

type Food struct {
	Id          string `json:"id"`
	FoodName    string `json:"foodName"`
	Description string `json:"description"`
}
type FoodIngredient struct {
	Id           string `json:"id"`
	FoodId       string `json:"foodId"`
	IngredientId string `json:"ingredientId"`
	Description  string `json:"description"`
}
type FoodImage struct {
	Id string `json:"id"`
	FoodId string `json:"foodId"`
	ImageId string `json:"imageId"`
}
type Ingredient struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type IngredientImage struct {
	Id string `json:"id"`
	IngredientId string `json:"ingredientId"`
	ImageId string `json:"imageId"`
}