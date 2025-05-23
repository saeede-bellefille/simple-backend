package dto

type RegisdterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Repeat   string `json:"repeat_password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}
