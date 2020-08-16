package subscription

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"

)

const PayTypeURL = api.BASE_URL + "paytype/ "

func CreatePayType(users user.Users) (user.PayType, error) {
	entity := user.PayType{}
	resp, _ := api.Rest().SetBody(users).Post(PayTypeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPayType(id string) (user.PayType, error) {
	entity := user.PayType{}
	resp, _ := api.Rest().Get(PayTypeURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeletePayType(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(PayTypeURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadPayTypeAll() ([]user.PayType, error) {
	entity := []user.PayType{}
	resp, _ := api.Rest().Get(PayTypeURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPayTypeAllOf() ([]user.PayType, error) {
	entity := []user.PayType{}
	resp, _ := api.Rest().Get(PayTypeURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
