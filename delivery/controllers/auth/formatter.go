package auth

type LoginRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

type LoginResponseFormat struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
