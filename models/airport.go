package models

import (
	"mime/multipart"
)

type Airport struct {
	Id           string  `json:"id"`
	Guid         string  `json:"guid"`
	Title        string  `json:"title"`
	CountryId    string  `json:"country_id"`
	CityId       string  `json:"city_id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Radius       string  `json:"radius"`
	Image        string  `json:"image"`
	Address      string  `json:"address"`
	TimezoneId   string  `json:"timezone_id"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	SearchText   string  `json:"search_text"`
	Code         string  `json:"code"`
	ProductCount int     `json:"product_count"`
	Gmt          string  `json:"gmt"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type UpAirport struct {
	Id           string  `json:"id"`
	Guid         string  `json:"guid"`
	Title        string  `json:"title"`
	CountryId    string  `json:"country_id"`
	CityId       string  `json:"city_id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Radius       string  `json:"radius"`
	Image        string  `json:"image"`
	Address      string  `json:"address"`
	TimezoneId   string  `json:"timezone_id"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	SearchText   string  `json:"search_text"`
	Code         string  `json:"code"`
	ProductCount int     `json:"product_count"`
	Gmt          string  `json:"gmt"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
type CreateAirport struct {
	Title        string  `json:"title"`
	Guid         string  `json:"guid"`
	CountryId    string  `json:"country_id"`
	CityId       string  `json:"city_id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Radius       string  `json:"radius"`
	Image        string  `json:"image"`
	Address      string  `json:"address"`
	TimezoneId   string  `json:"timezone_id"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	SearchText   string  `json:"search_text"`
	Code         string  `json:"code"`
	ProductCount int     `json:"product_count"`
	Gmt          string  `json:"gmt"`
}
type UploadAirport struct {
	Base string
	File multipart.FileHeader
}

type UpdateAirport struct {
	Id           string  `json:"id"`
	Guid         string  `json:"guid"`
	Title        string  `json:"title"`
	CountryId    string  `json:"country_id"`
	CityId       string  `json:"city_id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Radius       string  `json:"radius"`
	Image        string  `json:"image"`
	Address      string  `json:"address"`
	TimezoneId   string  `json:"timezone_id"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	SearchText   string  `json:"search_text"`
	Code         string  `json:"code"`
	ProductCount int     `json:"product_count"`
	Gmt          string  `json:"gmt"`
}

type AirportPrimaryKey struct {
	Id string `json:"id"`
}

type GetListAirportRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListAirportResponse struct {
	Count    int       `json:"count"`
	Airports []Airport `json:"airports"`
}
