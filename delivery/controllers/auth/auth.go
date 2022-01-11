package auth

import (
	"net/http"
	"time"
	"todos/delivery/common"
	"todos/repository/auth"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Repo auth.AuthInterface
}

func NewAuthControllers(aurepo auth.AuthInterface) *AuthController {
	return &AuthController{
		Repo: aurepo,
	}
}

func (authcon AuthController) LoginAuthCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}
		if err := c.Bind(&loginFormat); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		checkedUser, err := authcon.Repo.LoginUser(loginFormat.Name, loginFormat.Password)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		if err != nil || checkedUser.ID != 0 {
			if loginFormat.Name != "" && loginFormat.Password != "" {
				token, err := CreateTokenAuth(checkedUser.ID)
				if err != nil {
					return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
				}
				return c.JSON(
					http.StatusOK, map[string]interface{}{
						"message": "Successful Operation",
						"token":   token,
					},
				)
			}
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		} else {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

	}
}

func CreateTokenAuth(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("RAHASIA"))
}
