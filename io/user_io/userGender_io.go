package user_io

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"
)

const userGenderURL = api.BASE_URL + "user_gender/ "

func CreateUserGender(users user.UserGender) (user.UserGender, error) {
	entity := user.UserGender{}
	resp, _ := api.Rest().SetBody(users).Post(userGenderURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserGender(users user.UserGender) (user.UserGender, error) {
	entity := user.UserGender{}
	resp, _ := api.Rest().SetBody(users).Post(userGenderURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadUserGender(id string) (user.UserGender, error) {
	entity := user.UserGender{}
	resp, _ := api.Rest().Get(userGenderURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteUserGender(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(userGenderURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadUserGenderAll() ([]user.UserGender, error) {
	entity := []user.UserGender{}
	resp, _ := api.Rest().Get(userGenderURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserGenderAllOf() ([]user.UserGender, error) {
	entity := []user.UserGender{}
	resp, _ := api.Rest().Get(userGenderURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
