package handler

import (
	"templeTest/view/pages"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)
type ViewHandler struct {
	Component templ.Component
}

func (h ViewHandler) WallOfMemoriesView(c echo.Context) error {
	return render(c, pages.WallOfMemories())
}
