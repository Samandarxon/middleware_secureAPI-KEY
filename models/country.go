package models

import "mime/multipart"

type Country struct {
	Id        string `json:"id"`
	Guid      string `json:"guid"`
	Title     string `json:"title"`
	Code      string `json:"code"`
	Continent string `json:"continent"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpCountry struct {
	Id        string `json:"id"`
	Guid      string `json:"guid"`
	Title     string `json:"title"`
	Code      string `json:"code"`
	Continent string `json:"continent"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type UploadCountry struct {
	Base string
	File multipart.FileHeader
}
type CreateCountry struct {
	Guid      string `json:"guid"`
	Title     string `json:"title"`
	Code      string `json:"code"`
	Continent string `json:"continent"`
}

type UpdateCountry struct {
	Id        string `json:"id"`
	Guid      string `json:"guid"`
	Title     string `json:"title"`
	Code      string `json:"code"`
	Continent string `json:"continent"`
}

type CountryPrimaryKey struct {
	Id string `json:"id"`
}

type GetListCountryRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListCountryResponse struct {
	Count     int       `json:"count"`
	Countries []Country `json:"countries"`
}
