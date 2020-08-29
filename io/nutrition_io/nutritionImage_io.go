package nutrition_io

import (
	"afcmo/api"
	"afcmo/domain/nutrition"
	"errors"

)

const nutritionImageURL = api.BASE_URL + "nutrition/"

func CreateNutritionImage(image nutrition.NutritionImage) (nutrition.NutritionImage, error) {
	entity := nutrition.NutritionImage{}
	resp, _ := api.Rest().SetBody(image).Post(nutritionImageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutritionImage(id string) (nutrition.NutritionImage, error) {
	entity := nutrition.NutritionImage{}
	resp, _ := api.Rest().Get(nutritionImageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteNutritionImage(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(nutritionImageURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadNutritionImageAll() ([]nutrition.NutritionImage, error) {
	entity := []nutrition.NutritionImage{}
	resp, _ := api.Rest().Get(nutritionImageURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutritionImageAllOf() ([]nutrition.NutritionImage, error) {
	entity := []nutrition.NutritionImage{}
	resp, _ := api.Rest().Get(nutritionImageURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
