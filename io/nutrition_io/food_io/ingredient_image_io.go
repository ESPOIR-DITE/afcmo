package food_io

import (
	"afcmo/api"
	"afcmo/domain/nutrition"
	"errors"

)

const ingredientimageURL = api.BASE_URL + "ingredient_image/"

func CreateIngredientImage(nutritionObject nutrition.IngredientImage) (nutrition.IngredientImage, error) {
	entity := nutrition.IngredientImage{}
	resp, _ := api.Rest().SetBody(nutritionObject).Post(ingredientimageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadIngredientImage(id string) (nutrition.IngredientImage, error) {
	entity := nutrition.IngredientImage{}
	resp, _ := api.Rest().Get(ingredientimageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteIngredientImage(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(ingredientimageURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadIngredientImageAll() ([]nutrition.IngredientImage, error) {
	entity := []nutrition.IngredientImage{}
	resp, _ := api.Rest().Get(ingredientimageURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadIngredientImageAllOf(ingredientId string) ([]nutrition.IngredientImage, error) {
	entity := []nutrition.IngredientImage{}
	resp, _ := api.Rest().Get(ingredientimageURL + "readAllOf?id=" + ingredientId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadIngredientImageOf(ingredientId string) (nutrition.IngredientImage, error) {
	entity := nutrition.IngredientImage{}
	resp, _ := api.Rest().Get(ingredientimageURL + "readOf?id=" + ingredientId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
