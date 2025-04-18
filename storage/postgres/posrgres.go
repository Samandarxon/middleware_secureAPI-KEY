package postgres

import (
	"database/sql"
	"essy_travel/config"
	"essy_travel/storage"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	db      *sql.DB
	city    *CityRepo
	country *CountryRepo
	airport *AirportRepo
	upload  *UploadRepo
	user    *UserRepo
	login   *LoginRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {
	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) City() storage.CityRepoI {
	if s.city == nil {
		s.city = NewCityRepo(s.db)
	}
	return s.city
}

func (s *Store) Airport() storage.AirportRepoI {
	if s.airport == nil {
		s.airport = NewAirportRepo(s.db)
	}
	return s.airport
}

func (s *Store) Country() storage.CountryRepoI {
	if s.country == nil {
		s.country = NewCountryRepo(s.db)
	}
	return s.country
}

func (s *Store) Upload() storage.UploadRepoI {
	if s.upload == nil {
		s.upload = NewUploadRepo(s.db)
	}
	return s.upload
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}
	return s.user
}
func (s *Store) Login() storage.LoginRepoI {
	if s.login == nil {
		s.login = NewLoginRepo(s.db)
	}
	return s.login
}
