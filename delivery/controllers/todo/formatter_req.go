package todo

type CreateToDoRequestFormat struct {
	ProjectID   uint   `json:"projectid" form:"projectid"`
	UserID      uint   `json:"UserID" form:"UserID"`
	Task        string `json:"task" form:"task"`
	Status      string `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
}

type PutToDoRequestFormat struct {
	ProjectID   uint   `json:"projectid" form:"projectid"`
	UserID      uint   `json:"UserID" form:"UserID"`
	Task        string `json:"task" form:"task"`
	Status      string `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
}
