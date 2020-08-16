package image

type Images struct {
	Id          string `json:"id"`
	Video       []byte `json:"video"`
	Description string `json:"description"`
}
type Video struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Video       []byte `json:"video"`
}
