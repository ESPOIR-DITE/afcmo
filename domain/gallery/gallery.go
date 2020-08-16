package gallery

type Gallery struct {
	Id          string `json:"id"`
	Gallery     string `json:"gallery"`
	Description string `json:"description"`
}
type GalleryImage struct {
	Id        string `json:"id"`
	Image     string `json:"image"`
	GalleryId string `json:"galleryId"`
}
