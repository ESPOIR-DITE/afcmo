package food_io

import (
	"afcmo/api"
	"afcmo/domain/nutrition"
	"errors"

)

const foodURL = api.BASE_URL + "food/"

func CreateFood(nutritionObject nutrition.Food) (nutrition.Food, error) {
	entity := nutrition.Food{}
	resp, _ := api.Rest().SetBody(nutritionObject).Post(foodURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadFood(id string) (nutrition.Food, error) {
	entity := nutrition.Food{}
	resp, _ := api.Rest().Get(foodURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteFood(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(foodURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadFoodAll() ([]nutrition.Food, error) {
	entity := []nutrition.Food{}
	resp, _ := api.Rest().Get(foodURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadFoodAllOf() ([]nutrition.Food, error) {
	entity := []nutrition.Food{}
	resp, _ := api.Rest().Get(foodURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
