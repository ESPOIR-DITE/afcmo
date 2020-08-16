package user_io

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"

)

const accountURL = api.BASE_URL + "account/ "

func CreateAccount(users user.Account) (user.Account, error) {
	entity := user.Account{}
	resp, _ := api.Rest().SetBody(users).Post(accountURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAccount(id string) (user.Account, error) {
	entity := user.Account{}
	resp, _ := api.Rest().Get(accountURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteAccount(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(accountURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadAccountAll() ([]user.Account, error) {
	entity := []user.Account{}
	resp, _ := api.Rest().Get(accountURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAccountAllOf() ([]user.Account, error) {
	entity := []user.Account{}
	resp, _ := api.Rest().Get(accountURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
