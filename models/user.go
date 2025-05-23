package models

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type GetListUserRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListUserResponse struct {
	Count int    `json:"count"`
	Users []User `json:"users"`
}
