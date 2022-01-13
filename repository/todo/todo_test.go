package todo

import (
	"fmt"
	"testing"
	"todos/configs"
	"todos/entities"
	"todos/repository/user"
	"todos/utils"

	"github.com/stretchr/testify/assert"
)

func TestTodoRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	var newUser entities.User
	newUser.Name = "TestName1"
	newUser.Password = "TestPassword1"

	userRep := user.NewUsersRepo(db)
	_, err := userRep.Create(newUser)
	if err != nil {
		fmt.Println(err)
	}

	db.Migrator().DropTable(&entities.ToDo{})
	db.AutoMigrate(&entities.ToDo{})

	todoRepo := NewToDoRepo(db)

	t.Run(
		"Insert ToDo into Database", func(t *testing.T) {
			var mockInsertToDo entities.ToDo
			mockInsertToDo.UserID = 1
			mockInsertToDo.Task = "Lanjut project"
			mockInsertToDo.Status = "in Progress"
			mockInsertToDo.Description = "Nyelesaiin project"

			res, err := todoRepo.Create(mockInsertToDo)
			assert.Nil(t, err)
			assert.Equal(t, mockInsertToDo.Task, res.Task)
			assert.Equal(t, 1, int(res.ID))
		},
	)
	t.Run(
		"Select ToDo from Database where ToDoID 1", func(t *testing.T) {
			res, err := todoRepo.Get(1, 1)
			assert.Nil(t, err)
			assert.Equal(t, res, res)
		},
	)
	t.Run(
		"Select ToDo from Database where userID 1", func(t *testing.T) {
			res, err := todoRepo.GetAll(1)
			assert.Nil(t, err)
			assert.Equal(t, res, res)
		},
	)

	t.Run(
		"Update ToDo by ToDoID", func(t *testing.T) {
			var mockUpdateToDo entities.ToDo
			mockUpdateToDo.Task = "UPDATE Lanjut project"
			mockUpdateToDo.Status = "Reopen"
			mockUpdateToDo.Description = "UPDATE Nyelesaiin project"

			res, err := todoRepo.Update(mockUpdateToDo, 1, 1)
			assert.Nil(t, err)
			assert.Equal(t, res, res)
		},
	)

	t.Run(
		"Delete ToDo by ToDo Id", func(t *testing.T) {
			res, err := todoRepo.Delete(1, 1)
			assert.Nil(t, err)
			assert.Equal(t, res, res)
		},
	)
}
