package project

import (
	"errors"
	"todos/entities"
)

type mockProjectRepository struct{}

func (mpr mockProjectRepository) GetAll() ([]entities.Project, error) {
	return []entities.Project{
		{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}},
	}, nil
}
func (mpr mockProjectRepository) Get(userId int) (entities.Project, error) {
	return entities.Project{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, nil
}
func (mpr mockProjectRepository) Create(newUser entities.Project) (entities.Project, error) {
	return entities.Project{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, nil
}
func (mpr mockProjectRepository) Update(updateUser entities.Project, userId int) (entities.Project, error) {
	return entities.Project{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, nil
}
func (mpr mockProjectRepository) Delete(userId int) (entities.Project, error) {
	return entities.Project{ID: 1, ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, nil
}

type mockFalseProjectRepository struct{}

func (mpr mockFalseProjectRepository) GetAll() ([]entities.Project, error) {
	return []entities.Project{
		{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}},
	}, errors.New("Bad Request")
}
func (mpr mockFalseProjectRepository) Get(userId int) (entities.Project, error) {
	return entities.Project{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, errors.New("Bad Request")
}
func (mpr mockFalseProjectRepository) Create(newUser entities.Project) (entities.Project, error) {
	return entities.Project{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, errors.New("Bad Request")
}
func (mpr mockFalseProjectRepository) Update(updateUser entities.Project, userId int) (entities.Project, error) {
	return entities.Project{ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, errors.New("Bad Request")
}
func (mpr mockFalseProjectRepository) Delete(userId int) (entities.Project, error) {
	return entities.Project{ID: 1, ProjectName: "ProjectName1", UserId: 1, Todo: []entities.ToDo{}}, errors.New("Bad Request")
}
