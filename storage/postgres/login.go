package postgres

import (
	"database/sql"
	"essy_travel/models"
	"fmt"
)

type LoginRepo struct {
	db *sql.DB
}

func NewLoginRepo(db *sql.DB) *LoginRepo {
	return &LoginRepo{
		db: db,
	}
}

func (c *LoginRepo) GetByLogin(req models.Login) (*models.LoginRespons, error) {

	var (
		user  = models.LoginRespons{}
		query = `
		SELECT 
				"id",
				"name",
				"login",
				"password",
				"created_at",
				"updated_at"
		FROM "user" WHERE login=$1 AND password=$2`
	)
	var (
		Id        sql.NullString
		Name      sql.NullString
		Login     sql.NullString
		Password  sql.NullString
		CreatedAt sql.NullString
		UpdatedAt sql.NullString
	)
	fmt.Println(query)
	resp := c.db.QueryRow(query, req.Login, req.Password)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&Name,
		&Login,
		&Password,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	user = models.LoginRespons{
		User: models.User{
			Id:        Id.String,
			Name:      Name.String,
			Login:     Login.String,
			Password:  Password.String,
			CreatedAt: CreatedAt.String,
			UpdatedAt: UpdatedAt.String,
		},
	}
	return &user, nil
}
