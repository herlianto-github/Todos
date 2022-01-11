package project

import (
	"net/http"
	"todos/delivery/common"
	"todos/entities"
	"todos/repository/project"

	"github.com/labstack/echo/v4"
)

type ProjectsController struct {
	Repo project.ProjectInterface
}

func NewProjectsControllers(prrep project.ProjectInterface) *ProjectsController {
	return &ProjectsController{Repo: prrep}
}

// POST /projects/register
func (prcon ProjectsController) PostToDoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newProjectReq := CreateProjectRequestFormat{}

		if err := c.Bind(&newProjectReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newProject := entities.Project{
			ProjectName: newProjectReq.ProjectName,
			Todo:        newProjectReq.Todo,
		}

		_, err := prcon.Repo.Create(newProject)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}

// GET /projects
func (prcon ProjectsController) GetProjectsCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		projects, err := prcon.Repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		response := GetProjectsResponseFormat{
			Message: "Successful Operation",
			Data:    projects,
		}

		return c.JSON(http.StatusOK, response)
	}
}
