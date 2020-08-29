package food_io

import (
	"afcmo/api"
	"afcmo/domain/nutrition"
	"errors"

)

const foodIngredientURL = api.BASE_URL + "nutrition/"

func CreateFoodIngredient(nutritionObject nutrition.FoodIngredient) (nutrition.FoodIngredient, error) {
	entity := nutrition.FoodIngredient{}
	resp, _ := api.Rest().SetBody(nutritionObject).Post(foodIngredientURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateFoodIngredient(nutritionObject nutrition.FoodIngredient) (nutrition.FoodIngredient, error) {
	entity := nutrition.FoodIngredient{}
	resp, _ := api.Rest().SetBody(nutritionObject).Post(foodIngredientURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadFoodIngredient(id string) (nutrition.FoodIngredient, error) {
	entity := nutrition.FoodIngredient{}
	resp, _ := api.Rest().Get(foodIngredientURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteFoodIngredient(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(foodIngredientURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadFoodIngredientAll() ([]nutrition.FoodIngredient, error) {
	entity := []nutrition.FoodIngredient{}
	resp, _ := api.Rest().Get(foodIngredientURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadFoodIngredientAllOf() ([]nutrition.FoodIngredient, error) {
	entity := []nutrition.FoodIngredient{}
	resp, _ := api.Rest().Get(foodIngredientURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
