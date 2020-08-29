package nutrition_io

import (
	"afcmo/api"
	"afcmo/domain/nutrition"
	"errors"

)

const nutritionUserPlanURL = api.BASE_URL + "nutrition/"

func CreateNutritionUserPlan(nutritionObject nutrition.NutritionUserPlan) (nutrition.NutritionUserPlan, error) {
	entity := nutrition.NutritionUserPlan{}
	resp, _ := api.Rest().SetBody(nutritionObject).Post(nutritionUserPlanURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutritionUserPlan(id string) (nutrition.NutritionUserPlan, error) {
	entity := nutrition.NutritionUserPlan{}
	resp, _ := api.Rest().Get(nutritionUserPlanURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteNutritionUserPlan(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(nutritionUserPlanURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadNutritionUserPlanAll() ([]nutrition.NutritionUserPlan, error) {
	entity := []nutrition.NutritionUserPlan{}
	resp, _ := api.Rest().Get(nutritionUserPlanURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutritionUserPlanAllOf() ([]nutrition.NutritionUserPlan, error) {
	entity := []nutrition.NutritionUserPlan{}
	resp, _ := api.Rest().Get(nutritionUserPlanURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
