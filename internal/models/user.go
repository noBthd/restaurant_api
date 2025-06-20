package models

type User struct {
	ID       		int    `json:"id"`
	Email 	 		string `json:"email"`
	Password 		string `json:"password"`
	Creation_date 	string `json:"creation_date"`
	Is_admin 		bool   `json:"is_admin"`
}