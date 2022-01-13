package todo

type CreateToDoRequestFormat struct {
	Task        string `json:"task" form:"task"`
	Description string `json:"description" form:"description"`
	UserID      uint   `json:"userid" form:"userid"`
	ProjectID   *uint  `json:"projectid" form:"projectid"`
}

type PutToDoRequestFormat struct {
	ToDoID      int    `json:"todoid" form:"todoid"`
	Task        string `json:"task" form:"task"`
	Description string `json:"description" form:"description"`
}
