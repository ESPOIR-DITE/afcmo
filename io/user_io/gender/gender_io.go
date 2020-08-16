package gender

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"
)

const genderURL = api.BASE_URL + "gender/"

func CreateGender(users user.Gender) (user.Gender, error) {
	entity := user.Gender{}
	resp, _ := api.Rest().SetBody(users).Post(genderURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateGender(users user.Gender) (user.Gender, error) {
	entity := user.Gender{}
	resp, _ := api.Rest().SetBody(users).Post(genderURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadGender(id string) (user.Gender, error) {
	entity := user.Gender{}
	resp, _ := api.Rest().Get(genderURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteGender(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(genderURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadGenderAll() ([]user.Gender, error) {
	entity := []user.Gender{}
	resp, _ := api.Rest().Get(genderURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGenderAllOf() ([]user.Gender, error) {
	entity := []user.Gender{}
	resp, _ := api.Rest().Get(genderURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
