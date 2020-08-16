package health_io

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"
)

const healthURL = api.BASE_URL + "health/ "

func CreateHealth(users user.Health) (user.Health, error) {
	entity := user.Health{}
	resp, _ := api.Rest().SetBody(users).Post(healthURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHealth(id string) (user.Health, error) {
	entity := user.Health{}
	resp, _ := api.Rest().Get(healthURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteHealth(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(healthURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadHealthAll() ([]user.Health, error) {
	entity := []user.Health{}
	resp, _ := api.Rest().Get(healthURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadHealthAllOf() ([]user.Health, error) {
	entity := []user.Health{}
	resp, _ := api.Rest().Get(healthURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
