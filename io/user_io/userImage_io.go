package user_io

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"

)

const userImageURL = api.BASE_URL + "user/ "

func CreateUserImage(users user.UserImage) (user.UserImage, error) {
	entity := user.UserImage{}
	resp, _ := api.Rest().SetBody(users).Post(userImageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserImage(users user.UserImage) (user.UserImage, error) {
	entity := user.UserImage{}
	resp, _ := api.Rest().SetBody(users).Post(userImageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadUserImage(id string) (user.UserImage, error) {
	entity := user.UserImage{}
	resp, _ := api.Rest().Get(userImageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteUserImage(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(userImageURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadUserImageAll() ([]user.UserImage, error) {
	entity := []user.UserImage{}
	resp, _ := api.Rest().Get(userImageURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserImageAllOf() ([]user.UserImage, error) {
	entity := []user.UserImage{}
	resp, _ := api.Rest().Get(userImageURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
