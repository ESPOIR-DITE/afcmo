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
type Ingredient struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
