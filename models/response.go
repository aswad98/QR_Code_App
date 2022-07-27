package models

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type UPlodFormResponse struct {
	Name   *string `json:"name" db:"name"`
	Image  *[]byte `json:"image" db:"image"`
	QRCode *[]byte `json:"qr_code" db:"qr_code"`
}

func NewUPlodFormResponse() *UPlodFormResponse {
	return &UPlodFormResponse{
		Name:   new(string),
		Image:  new([]byte),
		QRCode: new([]byte),
	}
}

type GetAllFormResponse struct {
	ID     *int    `json:"id" db:"id"`
	Name   *string `json:"name" db:"name"`
	Image  *[]byte `json:"image" db:"image"`
	QRCode *[]byte `json:"qr_code" db:"qr_code"`
}

func NewGetAllFormResponse() *GetAllFormResponse {
	return &GetAllFormResponse{
		ID:     new(int),
		Name:   new(string),
		Image:  new([]byte),
		QRCode: new([]byte),
	}
}
