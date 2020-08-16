package gallery_io

import (
	"afcmo/api"
	"afcmo/domain/gallery"
	"errors"

)

const galleryImageURL = api.BASE_URL + "gallery_image/ "

func CreateGalleryImage(galleryImageObejct gallery.GalleryImage) (gallery.GalleryImage, error) {
	entity := gallery.GalleryImage{}
	resp, _ := api.Rest().SetBody(galleryImageObejct).Post(galleryImageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGalleryImage(id string) (gallery.GalleryImage, error) {
	entity := gallery.GalleryImage{}
	resp, _ := api.Rest().Get(galleryImageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteGalleryImage(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(galleryImageURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadGalleryImageAll() ([]gallery.GalleryImage, error) {
	entity := []gallery.GalleryImage{}
	resp, _ := api.Rest().Get(galleryImageURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGalleryImageAllOf() ([]gallery.GalleryImage, error) {
	entity := []gallery.GalleryImage{}
	resp, _ := api.Rest().Get(galleryImageURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
