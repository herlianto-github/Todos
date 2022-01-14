package project

import (
	"net/http"
	"todos/delivery/common"
	"todos/entities"
	"todos/repository/project"

	"github.com/golang-jwt/jwt"
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

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		if newProjectReq.ProjectName != "" {
			newProject := entities.Project{
				ProjectName: newProjectReq.ProjectName,
				UserId:      uint(userID),
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

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		projects, err := prcon.Repo.GetAll(userID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "Successful Operation",
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

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		projects, err := tdcon.Repo.Get(ProjectId.ProjectID, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "Successful Operation",
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

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		_, err := tdcon.Repo.Delete(ProjectId.ProjectID, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "Successful Operation",
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

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		newProject := entities.Project{
			ProjectName: PutProjectReq.ProjectName,
		}

		_, err := tdcon.Repo.Update(newProject, PutProjectReq.ProjectID, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}
