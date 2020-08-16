package nutrition_io

import (
	"afcmo/api"
	"afcmo/domain/nutrition"
	"errors"

)

const nutritionTypeURL = api.BASE_URL + "nutrition/ "

func CreateNutritionType(nutritionType nutrition.NutritionType) (nutrition.NutritionType, error) {
	entity := nutrition.NutritionType{}
	resp, _ := api.Rest().SetBody(nutritionType).Post(nutritionTypeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutritionType(id string) (nutrition.NutritionType, error) {
	entity := nutrition.NutritionType{}
	resp, _ := api.Rest().Get(nutritionTypeURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteNutritionType(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(nutritionTypeURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadNutritionTypeAll() ([]nutrition.NutritionType, error) {
	entity := []nutrition.NutritionType{}
	resp, _ := api.Rest().Get(nutritionTypeURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutritionTypeAllOf() ([]nutrition.NutritionType, error) {
	entity := []nutrition.NutritionType{}
	resp, _ := api.Rest().Get(nutritionTypeURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
