package models

type Form struct {
	Name   *string `json:"name"`
	Image  *[]byte `json:"image"`
	QRCode *[]byte `json:"qr_code"`
}

func NewForm() *Form {
	return &Form{
		Name:   new(string),
		Image:  new([]byte),
		QRCode: new([]byte),
	}
}
