package project

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todos/configs"
	"todos/delivery/controllers/auth"
	"todos/entities"
	projectRepo "todos/repository/project"
	"todos/repository/user"
	"todos/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestProject(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&entities.Project{})
	db.AutoMigrate(&entities.Project{})
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	var dummyProject entities.Project
	dummyProject.ProjectName = "ProjectName1"
	dummyProject.UserId = 1
	dummyProject.Todo = []entities.ToDo{}

	var dummyUser entities.User
	dummyUser.Name = "TestName1"
	dummyUser.Password = "TestPassword1"

	useRep := user.NewUsersRepo(db)
	_, err := useRep.Create(dummyUser)
	if err != nil {
		fmt.Println(err)
	}

	pRep := projectRepo.NewProjectRepo(db)
	_, err = pRep.Create(dummyProject)
	if err != nil {
		fmt.Println(err)
	}

	ec := echo.New()

	jwtToken := ""
	t.Run("POST /users/login", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "TestName1",
			"password": "TestPassword1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := ec.NewContext(req, res)
		context.SetPath("/users/login")

		authCon := auth.NewAuthControllers(mockAuthRepository{})
		authCon.LoginAuthCtrl()(context)

		responses := auth.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		fmt.Println(jwtToken)
		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)

	})
	t.Run(
		"Create Project", func(t *testing.T) {
			reqBody, _ := json.Marshal(
				map[string]interface{}{
					"projectname": "ProjectName1",
				},
			)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
			res := httptest.NewRecorder()
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
			req.Header.Set("Content-Type", "application/json")

			context := ec.NewContext(req, res)
			context.SetPath("/projects/")

			pCon := NewProjectsControllers(mockProjectRepository{})
			middleware.JWT([]byte("RAHASIA"))(pCon.PostProjectsCtrl())(context)

			responses := CreateProjectResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, "Successful Operation", responses.Message)
			assert.Equal(t, 200, res.Code)
		})
	t.Run(
		"Get All Project", func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
			req.Header.Set("Content-Type", "application/json")

			context := ec.NewContext(req, res)
			context.SetPath("/projects/")

			pCon := NewProjectsControllers(mockProjectRepository{})
			middleware.JWT([]byte("RAHASIA"))(pCon.GetAllProjectsCtrl())(context)

			responses := CreateProjectResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, "Successful Operation", responses.Message)
			assert.Equal(t, 200, res.Code)
		})
	t.Run(
		"Get Project", func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
			req.Header.Set("Content-Type", "application/json")

			context := ec.NewContext(req, res)
			context.SetPath("/projects/")
			context.SetParamNames("id")
			context.SetParamValues("1")

			pCon := NewProjectsControllers(mockProjectRepository{})
			middleware.JWT([]byte("RAHASIA"))(pCon.GetProjectsCtrl())(context)

			responses := CreateProjectResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, "Successful Operation", responses.Message)
			assert.Equal(t, 200, res.Code)
		})
	t.Run(
		"Update Project", func(t *testing.T) {
			requestbody, _ := json.Marshal(map[string]string{
				"projectname": "ProjectName1",
			})

			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestbody))
			res := httptest.NewRecorder()
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
			req.Header.Set("Content-Type", "application/json")

			context := ec.NewContext(req, res)
			context.SetPath("/projects/")
			context.SetParamNames("id")
			context.SetParamValues("1")

			pCon := NewProjectsControllers(mockProjectRepository{})
			middleware.JWT([]byte("RAHASIA"))(pCon.PutProjectsCtrl())(context)

			responses := CreateProjectResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, "Successful Operation", responses.Message)
			assert.Equal(t, 200, res.Code)
		})
	t.Run(
		"Delete Project", func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
			req.Header.Set("Content-Type", "application/json")

			req.Header.Set("Content-Type", "application/json")
			context := ec.NewContext(req, res)
			context.SetPath("/projects/")
			context.SetParamNames("id")
			context.SetParamValues("1")

			pCon := NewProjectsControllers(mockProjectRepository{})
			middleware.JWT([]byte("RAHASIA"))(pCon.DeleteProjectsCtrl())(context)

			responses := CreateProjectResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, "Successful Operation", responses.Message)
			assert.Equal(t, 200, res.Code)
		})
	//FALSE GET ALL INTERNAL SERVER ERROR
	t.Run("GET ALL PROJECT", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/projects/")

		pCon := NewProjectsControllers(mockFalseProjectRepository{})
		middleware.JWT([]byte("RAHASIA"))(pCon.GetAllProjectsCtrl())(context)
	})
}

type mockAuthRepository struct{}

func (mua mockAuthRepository) LoginUser(name, password string) (entities.User, error) {
	return entities.User{ID: 1, Name: "TestName1", Password: "TestPassword1"}, nil
}

type mockProjectRepository struct{}

func (mpr mockProjectRepository) GetAll(userId int) ([]entities.Project, error) {
	return []entities.Project{
		{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}},
	}, nil
}
func (mpr mockProjectRepository) Get(projectId, userId int) (entities.Project, error) {
	return entities.Project{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, nil
}
func (mpr mockProjectRepository) Create(newProject entities.Project) (entities.Project, error) {
	return entities.Project{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, nil
}
func (mpr mockProjectRepository) Update(updateProject entities.Project, projectId, userId int) (entities.Project, error) {
	return entities.Project{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, nil
}
func (mpr mockProjectRepository) Delete(projectid, userId int) (entities.Project, error) {
	return entities.Project{ID: 1, ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, nil
}

type mockFalseProjectRepository struct{}

func (mpr mockFalseProjectRepository) GetAll(userId int) ([]entities.Project, error) {
	return []entities.Project{
		{},
	}, errors.New("Bad Request")
}
func (mpr mockFalseProjectRepository) Get(projectId, userId int) (entities.Project, error) {
	return entities.Project{Todo: []entities.ToDo{}}, errors.New("Bad Request")
}
func (mpr mockFalseProjectRepository) Create(newUser entities.Project) (entities.Project, error) {
	return entities.Project{Todo: []entities.ToDo{}}, errors.New("Bad Request")
}
func (mpr mockFalseProjectRepository) Update(updateUser entities.Project, projectId, userId int) (entities.Project, error) {
	return entities.Project{Todo: []entities.ToDo{}}, errors.New("Bad Request")
}
func (mpr mockFalseProjectRepository) Delete(projectid, userId int) (entities.Project, error) {
	return entities.Project{Todo: []entities.ToDo{}}, errors.New("Bad Request")
}
