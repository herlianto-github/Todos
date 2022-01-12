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
func (prcon ProjectsController) PostProjectsCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newProjectReq := CreateProjectRequestFormat{}

		if err := c.Bind(&newProjectReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		if newProjectReq.ProjectName != "" && newProjectReq.UserID != 0 {
			newProject := entities.Project{
				ProjectName: newProjectReq.ProjectName,
				UserId:      newProjectReq.UserID,
			}

			_, err := prcon.Repo.Create(newProject)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
			}

			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		}

		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

}

// GET /projects
func (prcon ProjectsController) GetAllProjectsCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		userId := GetAllProjectRequestFormat{}

		if err := c.Bind(&userId); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		projects, err := prcon.Repo.GetAll(userId.UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "success",
				"data":    projects,
			},
		)
	}
}
func (tdcon ProjectsController) GetProjectsCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		ProjectId := GetProjectRequestFormat{}

		if err := c.Bind(&ProjectId); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		projects, err := tdcon.Repo.Get(ProjectId.ProjectID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "success",
				"data":    projects,
			},
		)
	}

}
func (tdcon ProjectsController) DeleteProjectsCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		ProjectId := DeleteProjectRequestFormat{}

		if err := c.Bind(&ProjectId); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		_, err := tdcon.Repo.Delete(ProjectId.ProjectID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "success",
			},
		)
	}

}
func (tdcon ProjectsController) PutProjectsCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		PutProjectReq := PutProjectRequestFormat{}

		if err := c.Bind(&PutProjectReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newProject := entities.Project{
			ProjectName: PutProjectReq.ProjectName,
		}

		_, err := tdcon.Repo.Update(newProject, PutProjectReq.ProjectID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}
