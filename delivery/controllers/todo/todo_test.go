package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todos/configs"
	"todos/entities"
	todoRepo "todos/repository/todo"
	"todos/utils"
)

func TestUsers(t *testing.T) {
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
					"userid":      1,
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
			res := httptest.NewRecorder()

			req.Header.Set("Content-Type", "application/json")
			context := ec.NewContext(req, res)
			context.SetPath("/todo/")

			tdCon := NewToDoControllers(mockTodoRepository{})
			tdCon.PostTodoCtrl()(context)

			responses := CreateToDoResponseFormat{}
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, "Successful Operation", responses.Message)
			assert.Equal(t, 200, res.Code)
		},
	)
	t.Run(
		"Get todo", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()

			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
			context := ec.NewContext(req, res)
			context.SetPath("/users")

			userCon := NewUsersControllers(mockUserRepository{})
			if err := middleware.JWT([]byte("RAHASIA"))(userCon.GetUsersCtrl())(context); err != nil {
				log.Fatal(err)
				return
			}

			var responses GetUsersResponseFormat
			json.Unmarshal([]byte(res.Body.Bytes()), &responses)
			assert.Equal(t, responses.Data[0].Name, "TestName1")
		},
	)

}

type mockTodoRepository struct{}

func (mtd mockTodoRepository) Create(todo entities.ToDo) (entities.ToDo, error) {
	return entities.ToDo{Task: "Lanjut project", Description: "Nyelesaiin project", UserID: 1}, nil
}
func (mur mockTodoRepository) GetAll(userId int) ([]entities.ToDo, error) {
	return []entities.ToDo{
		{Task: "Lanjut project", Description: "Lanjut projec", UserID: 1},
	}, nil
}
func (mur mockTodoRepository) Get(toDoId int) (entities.ToDo, error) {
	return entities.ToDo{Task: "Lanjut project", Description: "Lanjut projec", UserID: 1}, nil
}
func (mur mockTodoRepository) Update(newToDo entities.ToDo, toDoId int) (entities.ToDo, error) {
	return entities.ToDo{Task: "Lanjut project", Description: "Lanjut projec", UserID: 1}, nil
}
func (mur mockTodoRepository) Delete(toDoId int) (entities.ToDo, error) {
	return entities.ToDo{ID: 1}, nil
}
