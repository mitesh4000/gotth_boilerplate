package main

import (
	"templeTest/db"
	"templeTest/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	db.InitializeDb()
	UserHandler := handler.UserHandler{}
	shoppingListHandler := handler.ShoppingListHandler{}
	View := handler.ViewHandler{}

	app.Static("/static", "view")
	app.GET("/", UserHandler.HandleUserShow)
	app.GET("/update-button-text", UserHandler.HandleUserShow)
	app.GET("/shopping-list", shoppingListHandler.HandleShoppingListView)
	app.POST("/api/shopping-list", shoppingListHandler.CreateItem)
	app.GET("/api/shopping-list", shoppingListHandler.GetItems)
	app.GET("/api/shopping-list/:id", shoppingListHandler.GetItemByID)
	app.DELETE("/api/shopping-list/:id", shoppingListHandler.DeleteItem)
	app.PUT("/api/shopping-list/:id", shoppingListHandler.UpdateItem)
	app.PUT("/api/update-button", UserHandler.HandleUpdateButtonText)
	app.GET("/wall-of-memories", View.WallOfMemoriesView)
	app.Start(":3000")
}
