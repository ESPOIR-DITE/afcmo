package subscription

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"
)

const payDetailsURL = api.BASE_URL + "pay_details/ "

func CreatePayDetail(users user.PayDetail) (user.PayDetail, error) {
	entity := user.PayDetail{}
	resp, _ := api.Rest().SetBody(users).Post(payDetailsURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPayDetail(id string) (user.PayDetail, error) {
	entity := user.PayDetail{}
	resp, _ := api.Rest().Get(payDetailsURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeletePayDetail(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(payDetailsURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadPayDetailAll() ([]user.PayDetail, error) {
	entity := []user.PayDetail{}
	resp, _ := api.Rest().Get(payDetailsURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPayDetailAllOf() ([]user.PayDetail, error) {
	entity := []user.PayDetail{}
	resp, _ := api.Rest().Get(payDetailsURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
