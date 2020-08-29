package nutrition_io

import (
	"afcmo/api"
	"afcmo/domain/nutrition"
	"errors"

)

const nutritionURL = api.BASE_URL + "nutrition/"

func CreateNutrition(nutritionObject nutrition.Nutrition) (nutrition.Nutrition, error) {
	entity := nutrition.Nutrition{}
	resp, _ := api.Rest().SetBody(nutritionObject).Post(nutritionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutrition(id string) (nutrition.Nutrition, error) {
	entity := nutrition.Nutrition{}
	resp, _ := api.Rest().Get(nutritionURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteNutrition(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(nutritionURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadNutritionAll() ([]nutrition.Nutrition, error) {
	entity := []nutrition.Nutrition{}
	resp, _ := api.Rest().Get(nutritionURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutritionAllOf() ([]nutrition.Nutrition, error) {
	entity := []nutrition.Nutrition{}
	resp, _ := api.Rest().Get(nutritionURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
