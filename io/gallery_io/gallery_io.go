package gallery_io

import (
	"afcmo/api"
	"afcmo/domain/gallery"
	"errors"

)

const galleryURL = api.BASE_URL + "gallery/ "

func CreateGallery(galleryObejct gallery.Gallery) (gallery.Gallery, error) {
	entity := gallery.Gallery{}
	resp, _ := api.Rest().SetBody(galleryObejct).Post(galleryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGallery(id string) (gallery.Gallery, error) {
	entity := gallery.Gallery{}
	resp, _ := api.Rest().Get(galleryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteGallery(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(galleryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGalleryAll() ([]gallery.Gallery, error) {
	entity := []gallery.Gallery{}
	resp, _ := api.Rest().Get(galleryURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGalleryAllOf() ([]gallery.Gallery, error) {
	entity := []gallery.Gallery{}
	resp, _ := api.Rest().Get(galleryURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
