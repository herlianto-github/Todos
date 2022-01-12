package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todos/configs"
	"todos/entities"
	"todos/repository/user"
	"todos/utils"

	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
)

func TestUsers(t *testing.T) {

	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	var dummyUser entities.User
	dummyUser.Name = "TestName1"
	dummyUser.Password = "TestPassword1"

	useRep := user.NewUsersRepo(db)
	_, err := useRep.Create(dummyUser)
	if err != nil {
		fmt.Println(err)
	}
	ec := echo.New()

	t.Run("POST /users/register", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "TestName1",
			"password": "TestPassword1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := ec.NewContext(req, res)
		context.SetPath("/users/register")

		userCon := NewUsersControllers(mockUserRepository{})
		userCon.PostUserCtrl()(context)

		responses := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)

	})
}

type mockUserRepository struct{}

func (mur mockUserRepository) GetAll() ([]entities.User, error) {
	return []entities.User{
		{Name: "TestName1", Password: "TestPassword1"},
	}, nil
}
func (mur mockUserRepository) Get(userId int) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, nil
}
func (mur mockUserRepository) Create(newUser entities.User) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, nil
}
func (mur mockUserRepository) Update(updateUser entities.User, userId int) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, nil
}
func (mur mockUserRepository) Delete(userId int) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, nil
}
