package user_io

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"
)

const userURL = api.BASE_URL + "user/"

func CreateUser(users user.Users) (user.Users, error) {
	entity := user.Users{}
	resp, _ := api.Rest().SetBody(users).Post(userURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func updateUser(users user.Users) (user.Users, error) {
	entity := user.Users{}
	resp, _ := api.Rest().SetBody(users).Post(userURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadUser(id string) (user.Users, error) {
	entity := user.Users{}
	resp, _ := api.Rest().Get(userURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteUser(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(userURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadUserAll() ([]user.Users, error) {
	entity := []user.Users{}
	resp, _ := api.Rest().Get(userURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserAllOf() ([]user.Users, error) {
	entity := []user.Users{}
	resp, _ := api.Rest().Get(userURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
