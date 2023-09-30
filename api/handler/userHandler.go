package handler

import (
	"errors"
	"golang-api/api/models"
	"golang-api/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateItem creates a new user item
// curl -X POST http://localhost:8080/api/user -H "Content-Type: application/json" -d '{"name": "John Doe", "password": "password"}'
func CreateItem(c echo.Context) error {
	if db.DB == nil {
        return c.JSON(http.StatusInternalServerError, "Database not initialized")
    }

	user := new(models.User)
	
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	
	if user.Name == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name and password cannot be empty"})
	}

	// Check if a user with the same username and password already exists
	existingUser := &models.User{}
	if err := db.DB.Where("name = ? AND password = ?", user.Name, user.Password).First(existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		// If an existing user is found, return a conflict status code
		return c.JSON(http.StatusConflict, map[string]string{"error": "A user with the given username and password already exists"})
	}
	
	if err := db.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	
	return c.JSON(http.StatusCreated, user)
}

// GetItem retrieves an item by ID
// curl -X GET http://localhost:8080/api/user/{id} -H "Content-Type: application/json"
func GetItem(c echo.Context) error {
    // Check if the database is initialized
    if db.DB == nil {
        return c.JSON(http.StatusInternalServerError, "Database not initialized")
    }
    
    id := c.Param("id")
    var item models.User
    
    // Check if the item with the provided ID exists in the database
    if err := db.DB.First(&item, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Item not found")
    }
    
    // Return the found item as JSON
    return c.JSON(http.StatusOK, item)
}

// UpdateItem updates an existing item by ID
func UpdateItem(c echo.Context) error {
    if db.DB == nil {
        return c.JSON(http.StatusInternalServerError, "Database not initialized")
    }
    
    id := c.Param("id")
    var user models.User
    
    if err := db.DB.First(&user, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, "User not found")
    }
    
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    
    if err := db.DB.Save(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    
    return c.JSON(http.StatusOK, user)
}

// DeleteItem deletes an item by ID
func DeleteItem(c echo.Context) error {
    if db.DB == nil {
        return c.JSON(http.StatusInternalServerError, "Database not initialized")
    }
    
    id := c.Param("id")
    var user models.User
    
    if err := db.DB.First(&user, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, "User not found")
    }
    
    if err := db.DB.Delete(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    
    return c.NoContent(http.StatusNoContent)
}
