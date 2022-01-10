package users

import (
	"net/http"
	"todos/delivery/common"
	"todos/repository/users"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	Repo users.UsersInterface
}

func NewUsersControllers(usrep users.UsersInterface) *UsersController {
	return &UsersController{Repo: usrep}
}

// /users GET
func (uscon UsersController) GetUsersCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		users, err := uscon.Repo.Gets()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    users,
		})
	}
}
