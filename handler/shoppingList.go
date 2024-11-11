package handler

import (
	"fmt"
	"net/http"
	"templeTest/db"
	"templeTest/model"
	"templeTest/view/shoppingList"

	"github.com/labstack/echo/v4"
)

type ShoppingListHandler struct{}

func (h ShoppingListHandler) HandleShoppingListView(c echo.Context) error {
	return render(c, shoppingList.Page())
}

func (h ShoppingListHandler) CreateItem(c echo.Context) error {
	var newItem model.Item

	if err := c.Bind(&newItem); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"err": "invalid input"})
	}

	createdItem := db.DB.Create(newItem)
	return c.JSON(http.StatusOK, createdItem)
}

func (h ShoppingListHandler) GetItems(c echo.Context) error {
	var itemList []model.Item
	db.DB.Find(&itemList)
	return c.JSON(http.StatusOK, itemList)
}
func (h ShoppingListHandler) GetItemByID(c echo.Context) error {
	id := c.Param("id")
	var item model.Item

	if err := db.DB.First(&item, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"err": "item not found"})
	}
	return c.JSON(http.StatusOK, item)
}

func (h ShoppingListHandler) DeleteItem(c echo.Context) error {
	id := c.Param("id")
	var item model.Item
	if err := db.DB.Delete(&item, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"err": "item not found"})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h ShoppingListHandler) UpdateItem(c echo.Context) error {
	id := c.Param("id")
	var existingItem model.Item
	var updatedItem model.Item

	if err := db.DB.First(&existingItem, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"err": err})
	}

	if err := c.Bind(&updatedItem); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"err": "invalid input"})
	}
	fmt.Println(updatedItem)
	db.DB.Save(&updatedItem)
	return c.JSON(http.StatusOK, updatedItem)
}
