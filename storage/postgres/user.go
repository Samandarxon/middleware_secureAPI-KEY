package postgres

import (
	"database/sql"
	"essy_travel/models"
	"fmt"

	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (a *UserRepo) Create(req models.CreateUser) (*models.User, error) {
	var id = uuid.New().String()
	var query = `
	INSERT INTO "user" (
			"id",
			"name",
			"login",
			"password",
			"updated_at"
) VALUES ($1,$2,$3,$4,NOW())
`
	// fmt.Println(query, req)
	_, err := a.db.Exec(query,
		id,
		req.Name,
		req.Login,
		req.Password,
	)
	if err != nil {
		return &models.User{}, err
	}

	return a.GetById(models.UserPrimaryKey{Id: id})
}

func (c *UserRepo) GetById(req models.UserPrimaryKey) (*models.User, error) {

	var (
		user  = models.User{}
		query = `
		SELECT 
				"id",
				"name",
				"login",
				"password",
				"created_at",
				"updated_at"
		FROM "user" WHERE id=$1`
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
	resp := c.db.QueryRow(query, req.Id)
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
	user = models.User{
		Id:        Id.String,
		Name:      Name.String,
		Login:     Login.String,
		Password:  Password.String,
		CreatedAt: CreatedAt.String,
		UpdatedAt: UpdatedAt.String,
	}
	return &user, nil
}

func (c *UserRepo) GetList(req models.GetListUserRequest) (*models.GetListUserResponse, error) {
	var (
		users = models.GetListUserResponse{}

		Id        sql.NullString
		Name      sql.NullString
		Login     sql.NullString
		Password  sql.NullString
		CreatedAt sql.NullString
		UpdatedAt sql.NullString

		where  = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		query  = `
				SELECT 
				COUNT(*) OVER(),
				"id",
				"id",
				"name",
				"login",
				"password",
				"created_at",
				"updated_at"
				FROM "user"`
	)
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += where + offset + limit

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}
	fmt.Println(query)
	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&users.Count,
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
		user = models.User{
			Id:        Id.String,
			Name:      Name.String,
			Login:     Login.String,
			Password:  Password.String,
			CreatedAt: CreatedAt.String,
			UpdatedAt: UpdatedAt.String,
		}
		// fmt.Printf("%+v\n", user)
		users.Users = append(users.Users, user)
	}
	defer rows.Close()
	return &users, nil
}

func (c *UserRepo) Update(req models.UpdateUser) (*models.User, error) {
	var (
		query = `
			UPDATE "user" SET
				"name" = $2,
				"login" = $3,
				"password" = $4,
				updated_at  =  NOW() 
			WHERE id = $1`
	)
	_, err := c.db.Exec(query,
		req.Id,
		req.Name,
		req.Login,
		req.Password,
	)
	if err != nil {
		return nil, err
	}

	return c.GetById(models.UserPrimaryKey{Id: req.Id})
}

func (c *UserRepo) Delete(req models.UserPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM user WHERE id = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}
