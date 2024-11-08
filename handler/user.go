package handler

import (
	"templeTest/model"
	"templeTest/view/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	userData := model.User{
		Email: "mike@tutanota.com",
	}
	return render(c, user.Show(userData))

}
