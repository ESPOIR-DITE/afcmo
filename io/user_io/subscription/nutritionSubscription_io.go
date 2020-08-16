package subscription

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"
)

const nutritionSubscriptionURL = api.BASE_URL + "nutrition_subscription/ "

func CreateNutritionSubscription(users user.NutritionSubscription) (user.NutritionSubscription, error) {
	entity := user.NutritionSubscription{}
	resp, _ := api.Rest().SetBody(users).Post(nutritionSubscriptionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadNutritionSubscription(id string) (user.NutritionSubscription, error) {
	entity := user.NutritionSubscription{}
	resp, _ := api.Rest().Get(nutritionSubscriptionURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteNutritionSubscription(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(nutritionSubscriptionURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadNutritionSubscriptionAll() ([]user.NutritionSubscription, error) {
	entity := []user.NutritionSubscription{}
	resp, _ := api.Rest().Get(nutritionSubscriptionURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadNutritionSubscriptionAllOf() ([]user.NutritionSubscription, error) {
	entity := []user.NutritionSubscription{}
	resp, _ := api.Rest().Get(nutritionSubscriptionURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
