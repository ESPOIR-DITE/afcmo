package image_io

import (
	"afcmo/api"
	"afcmo/domain/image"
	"errors"
)

const imageURL = api.BASE_URL + "image/"

func CreateImage(images image.Images) (image.Images, error) {
	entity := image.Images{}
	resp, _ := api.Rest().SetBody(images).Post(imageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateImage(images image.Images) (image.Images, error) {
	entity := image.Images{}
	resp, _ := api.Rest().SetBody(images).Post(imageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadImage(id string) (image.Images, error) {
	entity := image.Images{}
	resp, _ := api.Rest().Get(imageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteImage(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(imageURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadImageAll() ([]image.Images, error) {
	entity := []image.Images{}
	resp, _ := api.Rest().Get(imageURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadImageAllOf() ([]image.Images, error) {
	entity := []image.Images{}
	resp, _ := api.Rest().Get(imageURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
