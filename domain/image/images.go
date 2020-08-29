package image

type Images struct {
	Id          string `json:"id"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}
type ImageHelper struct {
	Id          string `json:"id"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
type Video struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Video       []byte `json:"video"`
}
type Gifs struct{
	Id string `json:"id"`
	Gif []byte `json:"gif"`
	Description string `json:"description"`
}
