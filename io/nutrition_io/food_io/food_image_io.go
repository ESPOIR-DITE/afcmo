package food_io

import (
	"afcmo/api"
	"afcmo/domain/nutrition"
	"errors"
)

const foodimageURL = api.BASE_URL + "food_image/"

func CreateFoodImage(nutritionObject nutrition.FoodImage) (nutrition.FoodImage, error) {
	entity := nutrition.FoodImage{}
	resp, _ := api.Rest().SetBody(nutritionObject).Post(foodimageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateFoodImage(nutritionObject nutrition.FoodImage) (nutrition.FoodImage, error) {
	entity := nutrition.FoodImage{}
	resp, _ := api.Rest().SetBody(nutritionObject).Post(foodimageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadFoodImage(id string) (nutrition.FoodImage, error) {
	entity := nutrition.FoodImage{}
	resp, _ := api.Rest().Get(foodimageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteFoodImage(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(foodimageURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadFoodImageAll() ([]nutrition.FoodImage, error) {
	entity := []nutrition.FoodImage{}
	resp, _ := api.Rest().Get(foodimageURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadFoodImageAllOf() ([]nutrition.FoodImage, error) {
	entity := []nutrition.FoodImage{}
	resp, _ := api.Rest().Get(foodimageURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
