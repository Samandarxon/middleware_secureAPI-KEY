package models

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRespons struct {
	User    `json:"user"`
	API_KEY string `json:"api_key"`
}
