package auth

type LoginRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginResponseFormat struct {
	Message string `json:"message"`
	token   string `json:"token"`
}
