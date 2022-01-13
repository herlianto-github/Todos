package project

import (
	"fmt"
	"testing"
	"todos/configs"
	"todos/entities"
	"todos/repository/user"
	"todos/utils"

	"github.com/stretchr/testify/assert"
)

func TestProjectRepo(t *testing.T) {
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

	db.Migrator().DropTable(&entities.Project{})
	db.AutoMigrate(&entities.Project{})

	projectRepo := NewProjectRepo(db)

	t.Run("Insert Project into Database", func(t *testing.T) {
		var mockInsertProject entities.Project
		mockInsertProject.ProjectName = "ProjectName1"
		mockInsertProject.UserId = 1
		mockInsertProject.Todo = []entities.ToDo{}

		res, err := projectRepo.Create(mockInsertProject)
		assert.Nil(t, err)
		assert.Equal(t, mockInsertProject.ProjectName, res.ProjectName)
		assert.Equal(t, 1, int(res.ID))
	})
	t.Run("Select Project from Database where projectID 1", func(t *testing.T) {
		res, err := projectRepo.Get(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Select Project from Database where userID 1", func(t *testing.T) {
		res, err := projectRepo.GetAll(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("Update Project by projectID", func(t *testing.T) {
		var mockUpdateProject entities.Project
		mockUpdateProject.ProjectName = "UPDATE ProjectName1"
		mockUpdateProject.Todo = []entities.ToDo{}

		res, err := projectRepo.Update(mockUpdateProject, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("Delete Project by projectID", func(t *testing.T) {
		res, err := projectRepo.Delete(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
}
