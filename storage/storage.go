package storage

import "essy_travel/models"

type StorageI interface {
	City() CityRepoI
	Airport() AirportRepoI
	Country() CountryRepoI
	Upload() UploadRepoI
	User() UserRepoI
	Login() LoginRepoI
}

type UserRepoI interface {
	Create(req models.CreateUser) (*models.User, error)
	Update(req models.UpdateUser) (*models.User, error)
	GetById(req models.UserPrimaryKey) (*models.User, error)
	GetList(req models.GetListUserRequest) (*models.GetListUserResponse, error)
	Delete(req models.UserPrimaryKey) (string, error)
}

type LoginRepoI interface {
	GetByLogin(req models.Login) (*models.LoginRespons, error)
}

type CountryRepoI interface {
	Create(req models.CreateCountry) (*models.Country, error)
	Update(req models.UpdateCountry) (*models.Country, error)
	Upload(req models.UploadCountry) error
	GetById(req models.CountryPrimaryKey) (*models.Country, error)
	GetList(req models.GetListCountryRequest) (*models.GetListCountryResponse, error)
	Delete(req models.CountryPrimaryKey) (string, error)
}

type CityRepoI interface {
	Create(req models.CreateCity) (*models.City, error)
	Update(req models.UpdateCity) (*models.City, error)
	Upload(req models.UploadCity) error
	GetById(req models.CityPrimaryKey) (*models.City, error)
	GetList(req models.GetListCityRequest) (*models.GetListCityResponse, error)
	Delete(req models.CityPrimaryKey) (string, error)
}

type AirportRepoI interface {
	Create(req models.CreateAirport) (*models.Airport, error)
	Update(req models.UpdateAirport) (*models.Airport, error)
	Upload(req models.UploadAirport) error
	GetById(req models.AirportPrimaryKey) (*models.Airport, error)
	GetList(req models.GetListAirportRequest) (*models.GetListAirportResponse, error)
	Delete(req models.AirportPrimaryKey) (string, error)
}

type UploadRepoI interface {
	Upload(req *models.UploadRequest) error
}
