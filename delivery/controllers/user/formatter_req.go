package user

type RegisterUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

type PutUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}
