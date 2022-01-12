package todo

type CreateToDoRequestFormat struct {
	Task        string `json:"task" form:"task"`
	Description string `json:"description" form:"description"`
	UserID      uint   `json:"userid" form:"userid"`
	ProjectID   *uint  `json:"projectid" form:"projectid"`
}
type GetAllToDoRequestFormat struct {
	UserID int `json:"userid" form:"userid"`
}

type GetToDoRequestFormat struct {
	ToDoID int `json:"todoid" form:"todoid"`
}
type DeleteToDoRequestFormat struct {
	ToDoID int `json:"TodoID" form:"TodoID"`
}
type PutToDoRequestFormat struct {
	ToDoID      int    `json:"todoid" form:"todoid"`
	Task        string `json:"task" form:"task"`
	Description string `json:"description" form:"description"`
}
