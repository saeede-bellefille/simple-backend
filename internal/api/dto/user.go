package dto

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Repeat   string `json:"repeat_password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRole struct {
	Role string `json:"role"`
}
