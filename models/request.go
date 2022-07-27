package models

type UplodFormRequest struct {
	Name  *string `json:"name" form:"name"`
	Image *[]byte `json:"image" form:"image"`
}

func NewUplodFormRequest() *UplodFormRequest {
	return &UplodFormRequest{
		Name:  new(string),
		Image: new([]byte),
	}
}
