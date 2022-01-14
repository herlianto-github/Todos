package todo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todos/configs"
	"todos/entities"
	todoRepo "todos/repository/todo"
	"todos/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestTodo(t *testing.T) {
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDIxMzA0MzgsInVzZXJpZCI6MX0.LEZnOdeQAE5kJlHy1PByOS3NjND34q6WYv38stFruXc"
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&entities.ToDo{})
	db.AutoMigrate(&entities.ToDo{})

	var dummyTodo entities.ToDo

	dummyTodo.Task = "Lanjut project"
	dummyTodo.Description = "Nyelesaiin project"
	dummyTodo.UserID = 1

	tdRep := todoRepo.NewToDoRepo(db)
	_, err := tdRep.Create(dummyTodo)
	if err != nil {
		fmt.Println(err)
	}

	ec := echo.New()

	t.Run(
		"Create Todo", func(t *testing.T) {
			reqBody, _ := json.Marshal(
				map[string]interface{}{
					"task":        "Lanjut project",
					"description": "Nyelesaiin project",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
			req.Header.Set(
				"Authorization",
				fmt.Sprintf("Bearer %v", jwtToken),
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := ec.NewContext(req, res)
			context.SetPath("/todo")

			tdCon := NewToDoControllers(mockTodoRepository{})
			if err := middleware.JWT([]byte("RAHASIA"))(tdCon.PostTodoCtrl())(context); err != nil {
				log.Fatal(err)
				return
			}

			responses := CreateToDoResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, "Successful Operation", responses.Message)
			assert.Equal(t, 200, res.Code)
		},
	)
	t.Run(
		"Get all todo", func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			req.Header.Set(
				"Authorization",
				fmt.Sprintf("Bearer %v", jwtToken),
			)

			context := ec.NewContext(req, res)
			context.SetPath("/todo/all")

			todoCon := NewToDoControllers(mockTodoRepository{})
			if err := middleware.JWT([]byte("RAHASIA"))(todoCon.GetAllTodoCtrl())(context); err != nil {
				log.Fatal(err)
				return
			}

			var responses GetAllToDoResponseFormat
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, responses.Data[0].Task, "Lanjut project")
		},
	)
	t.Run(
		"Get todo", func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			req.Header.Set(
				"Authorization",
				fmt.Sprintf("Bearer %v", jwtToken),
			)

			context := ec.NewContext(req, res)
			context.SetPath("/users/:id")
			context.SetParamNames("id")
			context.SetParamValues("1")

			todoCon := NewToDoControllers(mockTodoRepository{})
			if err := middleware.JWT([]byte("RAHASIA"))(todoCon.GetTodoCtrl())(context); err != nil {
				log.Fatal(err)
				return
			}

			var responses GetAllToDoResponseFormat
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, "succes", responses.Message)
			assert.Equal(t, 200, res.Code)
		},
	)
	t.Run(
		"update todo", func(t *testing.T) {
			reqBody, _ := json.Marshal(
				map[string]string{
					"task":        "Lanjut project",
					"description": "Lanjut projec",
				},
			)
			req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
			res := httptest.NewRecorder()

			req.Header.Set(
				"Authorization",
				fmt.Sprintf("Bearer %v", jwtToken),
			)
			req.Header.Set("Content-Type", "application/json")
			context := ec.NewContext(req, res)
			context.SetPath("/todo/:id")
			context.SetParamNames("id")
			context.SetParamValues("1")

			tdCon := NewToDoControllers(mockTodoRepository{})
			if err := middleware.JWT([]byte("RAHASIA"))(tdCon.PutTodoCtrl())(context); err != nil {
				log.Fatal(err)
				return
			}

			responses := PutToDoResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)

			assert.Equal(t, "Successful Operation", responses.Message)
			assert.Equal(t, 200, res.Code)
		},
	)
	t.Run(
		"Delete todo", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/", nil)
			res := httptest.NewRecorder()

			req.Header.Set(
				"Authorization",
				fmt.Sprintf("Bearer %v", jwtToken),
			)
			req.Header.Set("Content-Type", "application/json")
			context := ec.NewContext(req, res)
			context.SetPath("/users/:id")
			context.SetParamNames("id")
			context.SetParamValues("1")

			tdCon := NewToDoControllers(mockTodoRepository{})
			if err := middleware.JWT([]byte("RAHASIA"))(tdCon.DeleteTodoCtrl())(context); err != nil {
				log.Fatal(err)
				return
			}

			responses := DeleteToDoResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)

			assert.Equal(t, "success", responses.Message)
			assert.Equal(t, 200, res.Code)

		},
	)
}

func TestFalseUsers(t *testing.T) {
	e := echo.New()
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDIxMTEyNTMsInVzZXJpZCI6MX0.B4Gfcyq2G2eO0q-W5FbX2ecPEDwDxiAbEcBKnzF_NNo"
	t.Run(
		"Create todo", func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			res := httptest.NewRecorder()
			req.Header.Set(
				"Authorization",
				fmt.Sprintf("Bearer %v", jwtToken),
			)
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/todo/")

			todoCon := NewToDoControllers(mockFalseTodoRepository{})
			if err := middleware.JWT([]byte("RAHASIA"))(todoCon.PostTodoCtrl())(context); err != nil {
				log.Fatal(err)
				return
			}

			responses := CreateToDoResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, responses.Message, "Internal Server Error")
			assert.Equal(t, res.Code, 500)
		},
	)
	t.Run(
		"Create todo", func(t *testing.T) {

			reqBody, _ := json.Marshal(
				map[string]int{
					"task": 1,
				},
			)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
			res := httptest.NewRecorder()
			req.Header.Set(
				"Authorization",
				fmt.Sprintf("Bearer %v", jwtToken),
			)
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/todo/")

			todoCon := NewToDoControllers(mockFalseTodoRepository{})
			middleware.JWT([]byte("RAHASIA"))(todoCon.PostTodoCtrl())(context)

			responses := CreateToDoResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, responses.Message, "Bad Request")
			assert.Equal(t, res.Code, 400)
		},
	)

	t.Run(
		"GET All Todo", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			req.Header.Set(
				"Authorization",
				fmt.Sprintf("Bearer %v", jwtToken),
			)

			context := e.NewContext(req, res)
			context.SetPath("/todo")

			todoCon := NewToDoControllers(mockFalseTodoRepository{})
			middleware.JWT([]byte("RAHASIA"))(todoCon.PostTodoCtrl())(context)

			var responses GetAllToDoResponseFormat
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, responses.Message, "Internal Server Error")
			assert.Equal(t, res.Code, 500)
		},
	)
	t.Run(
		"GET Todo fail", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set(
				"Authorization",
				fmt.Sprintf("Bearer %v", jwtToken),
			)
			context := e.NewContext(req, res)
			context.SetPath("/todo/:id")

			todoCon := NewToDoControllers(mockFalseTodoRepository{})
			middleware.JWT([]byte("RAHASIA"))(todoCon.GetTodoCtrl())(context)

			var responses GetToDoResponseFormat
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, responses.Message, "Bad Request")
			assert.Equal(t, res.Code, 400)
		},
	)
	t.Run(
		"GET Todo fail", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set(
				"Authorization",
				fmt.Sprintf("Bearer %v", jwtToken),
			)
			context := e.NewContext(req, res)
			context.SetPath("/todo/:id")
			context.SetParamNames("id")
			context.SetParamValues("2")

			todoCon := NewToDoControllers(mockFalseTodoRepository{})
			err := middleware.JWT([]byte("RAHASIA"))(todoCon.GetTodoCtrl())(context)
			if err != nil {
				log.Error(err)
			}

			var responses GetToDoResponseFormat
			err = json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			if err != nil {
				log.Error(err)
			}

			assert.Equal(t, responses.Message, "Task not found")

		},
	)
	t.Run(
		"PUT /todos/", func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPut, "/", nil)
			res := httptest.NewRecorder()

			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/todos/")

			todoCon := NewToDoControllers(mockFalseTodoRepository{})
			middleware.JWT([]byte("RAHASIA"))(todoCon.PutTodoCtrl())(context)

			responses := PutToDoResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)

			assert.Equal(t, responses.Message, "Bad Request")
			assert.Equal(t, res.Code, 400)
		},
	)
	//t.Run(
	//	"PUT /users/:id", func(t *testing.T) {
	//		reqBody, _ := json.Marshal(
	//			map[string]int{
	//				"name": 1,
	//			},
	//		)
	//
	//		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
	//		res := httptest.NewRecorder()
	//
	//		req.Header.Set("Content-Type", "application/json")
	//		context := e.NewContext(req, res)
	//		context.SetPath("/users/:id")
	//		context.SetParamNames("id")
	//		context.SetParamValues("1")
	//
	//		userCon := NewUsersControllers(mockFalseUserRepository{})
	//		userCon.EditUserCtrl()(context)
	//
	//		var responses GetUserResponseFormat
	//		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
	//		assert.Equal(t, responses.Message, "Bad Request")
	//		assert.Equal(t, res.Code, 400)
	//	},
	//)
	//t.Run(
	//	"PUT /users/:id", func(t *testing.T) {
	//		reqBody, _ := json.Marshal(
	//			map[string]string{
	//				"name": "TestName1",
	//			},
	//		)
	//
	//		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
	//		res := httptest.NewRecorder()
	//
	//		req.Header.Set("Content-Type", "application/json")
	//		context := e.NewContext(req, res)
	//		context.SetPath("/users/:id")
	//		context.SetParamNames("id")
	//		context.SetParamValues("2")
	//
	//		userCon := NewUsersControllers(mockFalseUserRepository{})
	//		userCon.EditUserCtrl()(context)
	//
	//		var responses GetUserResponseFormat
	//		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
	//		assert.Equal(t, responses.Message, "Not Found")
	//		assert.Equal(t, res.Code, 404)
	//	},
	//)
	//t.Run(
	//	"DELETE /users/:id", func(t *testing.T) {
	//
	//		req := httptest.NewRequest(http.MethodDelete, "/", nil)
	//		res := httptest.NewRecorder()
	//
	//		req.Header.Set("Content-Type", "application/json")
	//		context := e.NewContext(req, res)
	//		context.SetPath("/users/:id")
	//
	//		userCon := NewUsersControllers(mockFalseUserRepository{})
	//		userCon.DeleteUserCtrl()(context)
	//
	//		responses := DeleteUserResponseFormat{}
	//		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
	//
	//		assert.Equal(t, responses.Message, "Bad Request")
	//		assert.Equal(t, res.Code, 400)
	//	},
	//)
	//t.Run(
	//	"DELETE /users/:id", func(t *testing.T) {
	//
	//		req := httptest.NewRequest(http.MethodDelete, "/", nil)
	//		res := httptest.NewRecorder()
	//
	//		req.Header.Set("Content-Type", "application/json")
	//		context := e.NewContext(req, res)
	//		context.SetPath("/users/:id")
	//		context.SetParamNames("id")
	//		context.SetParamValues("1")
	//
	//		userCon := NewUsersControllers(mockFalseUserRepository{})
	//		userCon.DeleteUserCtrl()(context)
	//
	//		responses := DeleteUserResponseFormat{}
	//		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
	//
	//		assert.Equal(t, responses.Message, "Not Found")
	//		assert.Equal(t, res.Code, 404)
	//	},
	//)

}

type mockTodoRepository struct{}

func (mtd mockTodoRepository) Create(todo entities.ToDo) (entities.ToDo, error) {
	return entities.ToDo{Task: "Lanjut project", Description: "Nyelesaiin project", UserID: 1}, nil
}
func (mtd mockTodoRepository) GetAll(userId int) ([]entities.ToDo, error) {
	return []entities.ToDo{
		{ID: 1, Task: "Lanjut project", Description: "Lanjut projec", UserID: 1},
	}, nil
}
func (mtd mockTodoRepository) Get(toDoId, userid int) ([]entities.ToDo, error) {
	return []entities.ToDo{{ID: 2, Task: "Lanjut project", Description: "Lanjut projec", UserID: 1}}, nil
}
func (mtd mockTodoRepository) Update(newToDo entities.ToDo, toDoId, userID int) ([]entities.ToDo, error) {
	return []entities.ToDo{{Task: "Lanjut project", Description: "Lanjut projec", UserID: 1}}, nil
}
func (mtd mockTodoRepository) Delete(toDoId, userid int) (entities.ToDo, error) {
	return entities.ToDo{ID: 1}, nil
}

type mockFalseTodoRepository struct{}

func (mtd mockFalseTodoRepository) GetAll(userId int) ([]entities.ToDo, error) {
	return []entities.ToDo{
		{Task: "Lanjut project", Description: "Nyelesaiin project", UserID: 1},
	}, errors.New("Bad Request")
}
func (mtd mockFalseTodoRepository) Get(toDoId, userId int) ([]entities.ToDo, error) {
	return []entities.ToDo{{ID: 1, Task: "Istirahat"}}, errors.New("Bad Request")
}
func (mtd mockFalseTodoRepository) Create(todo entities.ToDo) (entities.ToDo, error) {
	return entities.ToDo{
		Task: "Lanjut project", Description: "Nyelesaiin project", UserID: 1,
	}, errors.New("Bad Request")
}
func (mtd mockFalseTodoRepository) Delete(toDoId, userId int) (entities.ToDo, error) {
	return entities.ToDo{ID: 1}, errors.New("Bad Request")
}
func (mtd mockFalseTodoRepository) Update(newToDo entities.ToDo, toDoId, userID int) ([]entities.ToDo, error) {
	return []entities.ToDo{{Task: "Lanjut project", Description: "Lanjut projec", UserID: 1}}, errors.New("Bad Request")
}
