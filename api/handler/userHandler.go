package handler

import (
	"golang-api/api/models"
	"golang-api/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateItem creates a new item
func CreateItem(c echo.Context) error {
    item := new(models.User)
    if err := c.Bind(item); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    
    db.DB.Create(item)
    
    return c.JSON(http.StatusCreated, item)
}

// GetItem retrieves an item by ID
func GetItem(c echo.Context) error {
    id := c.Param("id")
    var item models.User
    
    if err := db.DB.First(&item, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Item not found")
    }
    
    return c.JSON(http.StatusOK, item)
}

// UpdateItem updates an existing item by ID
func UpdateItem(c echo.Context) error {
    id := c.Param("id")
    var item models.User
    
    if err := db.DB.First(&item, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Item not found")
    }
    
    if err := c.Bind(&item); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    
    db.DB.Save(&item)
    
    return c.JSON(http.StatusOK, item)
}

// DeleteItem deletes an item by ID
func DeleteItem(c echo.Context) error {
    id := c.Param("id")
    var item models.User
    
    if err := db.DB.First(&item, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Item not found")
    }
    
    db.DB.Delete(&item)
    
    return c.NoContent(http.StatusNoContent)
}