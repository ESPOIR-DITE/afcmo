package food_io

import (
	"afcmo/api"
	"afcmo/domain/nutrition"
	"errors"

)

const ingredientURL = api.BASE_URL + "nutrition/ "

func CreateIngredient(nutritionObject nutrition.Ingredient) (nutrition.Ingredient, error) {
	entity := nutrition.Ingredient{}
	resp, _ := api.Rest().SetBody(nutritionObject).Post(ingredientURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadIngredient(id string) (nutrition.Ingredient, error) {
	entity := nutrition.Ingredient{}
	resp, _ := api.Rest().Get(ingredientURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteIngredient(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(ingredientURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadIngredientAll() ([]nutrition.Ingredient, error) {
	entity := []nutrition.Ingredient{}
	resp, _ := api.Rest().Get(ingredientURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadIngredientAllOf() ([]nutrition.Ingredient, error) {
	entity := []nutrition.Ingredient{}
	resp, _ := api.Rest().Get(ingredientURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
