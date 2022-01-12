package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"todos/configs"
	"todos/delivery/controllers/auth"
	"todos/entities"
	"todos/repository/user"
	"todos/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

		responses := LoginUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)

	})
	t.Run("GET /users", func(t *testing.T) {

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

	})
	t.Run("GET /users/:id", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := ec.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userCon := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userCon.GetUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		var responses GetUserResponseFormat
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, responses.Data.Name, responses.Data.Name)
	})
	t.Run("PUT /users/:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "TestName1",
			"password": "TestPassword1",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := ec.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userCon := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userCon.EditUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := PutUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)
	})
	t.Run("DELETE /users/:id", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := ec.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userCon := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userCon.DeleteUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := DeleteUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)
	})

}

func TestFalseUsers(t *testing.T) {
	e := echo.New()

	t.Run("POST /users/register", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/register")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.PostUserCtrl()(context)

		responses := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, responses.Message, "Internal Server Error")
		assert.Equal(t, res.Code, 500)
	})
	t.Run("POST /users/register", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]int{
			"name": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/register")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.PostUserCtrl()(context)

		responses := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, responses.Message, "Bad Request")
		assert.Equal(t, res.Code, 400)
	})
	t.Run("POST /users/login", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		authCon := auth.NewAuthControllers(mockFalseAuthRepository{})
		authCon.LoginAuthCtrl()(context)

		responses := LoginUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, responses.Message, "Internal Server Error")
		assert.Equal(t, res.Code, 500)
	})
	t.Run("GET /users", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.GetUsersCtrl()(context)

		var responses GetUsersResponseFormat
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, responses.Message, "Internal Server Error")
		assert.Equal(t, res.Code, 500)
	})
	t.Run("GET /users/:id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.GetUserCtrl()(context)

		var responses GetUsersResponseFormat
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, responses.Message, "Bad Request")
		assert.Equal(t, res.Code, 400)
	})
	t.Run("GET /users/:id", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.GetUserCtrl()(context)

		var responses GetUsersResponseFormat
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, responses.Message, "Not Found")
		assert.Equal(t, res.Code, 404)
	})
	t.Run("PUT /users/:id", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.EditUserCtrl()(context)

		responses := PutUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, responses.Message, "Bad Request")
		assert.Equal(t, res.Code, 400)
	})
	t.Run("PUT /users/:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"name": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.EditUserCtrl()(context)

		var responses GetUserResponseFormat
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, responses.Message, "Bad Request")
		assert.Equal(t, res.Code, 400)
	})
	t.Run("PUT /users/:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name": "TestName1",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.EditUserCtrl()(context)

		var responses GetUserResponseFormat
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, responses.Message, "Not Found")
		assert.Equal(t, res.Code, 404)
	})
	t.Run("DELETE /users/:id", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.DeleteUserCtrl()(context)

		responses := DeleteUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, responses.Message, "Bad Request")
		assert.Equal(t, res.Code, 400)
	})
	t.Run("DELETE /users/:id", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.DeleteUserCtrl()(context)

		responses := DeleteUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, responses.Message, "Not Found")
		assert.Equal(t, res.Code, 404)
	})

}

type mockAuthRepository struct{}

func (mua mockAuthRepository) LoginUser(name, password string) (entities.User, error) {
	return entities.User{ID: 1, Name: "TestName1", Password: "TestPassword1"}, nil
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
	return entities.User{ID: 1, Name: "TestName1", Password: "TestPassword1"}, nil
}

// FALSE SECTION
type mockFalseAuthRepository struct{}

func (mua mockFalseAuthRepository) LoginUser(name, password string) (entities.User, error) {
	return entities.User{ID: 1, Name: "TestName1", Password: "TestPassword1"}, errors.New("Bad Request")
}

type mockFalseUserRepository struct{}

func (mur mockFalseUserRepository) GetAll() ([]entities.User, error) {
	return []entities.User{
		{Name: "TestName1", Password: "TestPassword1"},
	}, errors.New("Bad Request")
}
func (mur mockFalseUserRepository) Get(userId int) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, errors.New("Bad Request")
}
func (mur mockFalseUserRepository) Create(newUser entities.User) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, errors.New("Bad Request")
}
func (mur mockFalseUserRepository) Update(updateUser entities.User, userId int) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, errors.New("Bad Request")
}
func (mur mockFalseUserRepository) Delete(userId int) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, errors.New("Bad Request")
}
