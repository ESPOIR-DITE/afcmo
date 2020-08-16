package health_io

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"
)

const userHealthURL = api.BASE_URL + "user/ "

func CreateUserHealth(users user.UserHealth) (user.UserHealth, error) {
	entity := user.UserHealth{}
	resp, _ := api.Rest().SetBody(users).Post(userHealthURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadUserHealth(id string) (user.UserHealth, error) {
	entity := user.UserHealth{}
	resp, _ := api.Rest().Get(userHealthURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteUserHealth(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(userHealthURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadUserHealthAll() ([]user.UserHealth, error) {
	entity := []user.UserHealth{}
	resp, _ := api.Rest().Get(userHealthURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserHealthAllOf() ([]user.UserHealth, error) {
	entity := []user.UserHealth{}
	resp, _ := api.Rest().Get(userHealthURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
