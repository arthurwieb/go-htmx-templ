package handler

import (
	"github.com/arthurwieb/got/model"
	"github.com/arthurwieb/got/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "arthur@me.com",
	}
	return render(c, user.Show(u))
}
