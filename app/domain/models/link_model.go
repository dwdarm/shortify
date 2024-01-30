package models

import "github.com/asaskevich/govalidator"

type Link struct {
	Slug   string `valid:"required" json:"slug" bson:"slug"`
	Href   string `valid:"url,required" json:"href" bson:"href"`
	QrCode string `valid:"optional" json:"qr_code" bson:"qr_code,omitempty"`
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
