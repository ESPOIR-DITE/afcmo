package image_io

import (
	"afcmo/api"
	"afcmo/domain/image"
	"errors"

)

const videoURL = api.BASE_URL + "video/"

func CreateVideo(images image.Video) (image.Video, error) {
	entity := image.Video{}
	resp, _ := api.Rest().SetBody(images).Post(videoURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateVideo(images image.Video) (image.Video, error) {
	entity := image.Video{}
	resp, _ := api.Rest().SetBody(images).Post(videoURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadVideo(id string) (image.Video, error) {
	entity := image.Video{}
	resp, _ := api.Rest().Get(videoURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteVideo(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(videoURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadVideoAll() ([]image.Video, error) {
	entity := []image.Video{}
	resp, _ := api.Rest().Get(videoURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadVideoAllOf() ([]image.Video, error) {
	entity := []image.Video{}
	resp, _ := api.Rest().Get(videoURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
