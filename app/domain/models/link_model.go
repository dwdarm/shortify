package models

import "github.com/asaskevich/govalidator"

type Link struct {
	Id     string `valid:"optional" json:"id"`
	Slug   string `valid:"required" json:"slug"`
	Href   string `valid:"url,required" json:"href"`
	QrCode string `valid:"optional" json:"qr_code"`
}

func NewLink(slug string, href string, qrCode string) (*Link, error) {
	link := &Link{
		Slug:   slug,
		Href:   href,
		QrCode: qrCode,
	}

	_, err := govalidator.ValidateStruct(link)
	if err != nil {
		return nil, err
	}

	return link, nil
}
