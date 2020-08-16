package nutrition_io

import (
	"afcmo/api"
	"afcmo/domain/nutrition"
	"errors"

)

const nutritionVideoURL = api.BASE_URL + "nutrition/ "

func CreateNutritionVideo(nutritionObject nutrition.NutritionVideo) (nutrition.NutritionVideo, error) {
	entity := nutrition.NutritionVideo{}
	resp, _ := api.Rest().SetBody(nutritionObject).Post(nutritionVideoURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutritionVideo(id string) (nutrition.NutritionVideo, error) {
	entity := nutrition.NutritionVideo{}
	resp, _ := api.Rest().Get(nutritionVideoURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteNutritionVideo(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(nutritionVideoURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadNutritionAllVideo() ([]nutrition.NutritionVideo, error) {
	entity := []nutrition.NutritionVideo{}
	resp, _ := api.Rest().Get(nutritionVideoURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutritionAllOfVideo() ([]nutrition.NutritionVideo, error) {
	entity := []nutrition.NutritionVideo{}
	resp, _ := api.Rest().Get(nutritionVideoURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
