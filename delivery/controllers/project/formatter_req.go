package project

type CreateProjectRequestFormat struct {
	ProjectName string `json:"projectname" form:"projectname"`
	UserID      uint   `json:"userid" form:"userid"`
}

type GetAllProjectRequestFormat struct {
	UserID int `json:"userid" form:"userid"`
}

type GetProjectRequestFormat struct {
	ProjectID int `json:"projectid" form:"projectid"`
}
type DeleteProjectRequestFormat struct {
	ProjectID int `json:"projectid" form:"TodoID"`
}
type PutProjectRequestFormat struct {
	ProjectName string `json:"projectname" form:"projectname"`
	ProjectID   int    `json:"projectid" form:"projectid"`
}
