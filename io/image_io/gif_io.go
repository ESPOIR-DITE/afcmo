package image_io

import (
	"afcmo/api"
	"afcmo/domain/image"
	"errors"
)

const gifURL = api.BASE_URL + "gifs/"

func CreateGif(images image.Gifs) (image.Gifs, error) {
	entity := image.Gifs{}
	resp, _ := api.Rest().SetBody(images).Post(gifURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateGif(images image.Gifs) (image.Gifs, error) {
	entity := image.Gifs{}
	resp, _ := api.Rest().SetBody(images).Post(gifURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGif(id string) (image.Gifs, error) {
	entity := image.Gifs{}
	resp, _ := api.Rest().Get(gifURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteGif(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(gifURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadGifAll() ([]image.Gifs, error) {
	entity := []image.Gifs{}
	resp, _ := api.Rest().Get(gifURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGifAllOf() ([]image.Gifs, error) {
	entity := []image.Gifs{}
	resp, _ := api.Rest().Get(gifURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
