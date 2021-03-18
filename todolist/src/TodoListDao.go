package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetTodoItems(c *gin.Context) error{
	var todoItems []TodoItem
	
	if err := DBHelper.Find(&todoItems).Error; err != nil {
		return UnknownError(err.Error())
	} else {
		c.JSON(http.StatusOK, todoItems)
		return nil
	}
}

func CreateTodoItem(c *gin.Context) error {
	todoItem := TodoItem{Status: 1, Created: time.Now()}

	if err := c.BindJSON(&todoItem); err != nil {
		return ParameterError(err.Error())
	} else if err := DBHelper.Create(&todoItem).Error; err != nil {
		return UnknownError(err.Error())
	} else {
		c.JSON(http.StatusOK, todoItem)
		return nil
	}
}

func UpdateTodoItem(c *gin.Context) error{
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return ParameterError(err.Error())
	}

	todoItem := TodoItem{Id: id,}
	updateTodoItem := TodoItem{}

	if err := c.BindJSON(&updateTodoItem); err != nil {
		return ParameterError(err.Error())
	}

	if err := DBHelper.First(&todoItem).Error; err != nil {
		return UnknownError(err.Error())
	}

	todoItem.Status = updateTodoItem.Status
	todoItem.Subject = updateTodoItem.Subject
	todoItem.Updated = time.Now()

	if err := DBHelper.Save(&todoItem).Error; err != nil {
		return UnknownError(err.Error())
	}else {
		c.JSON(http.StatusOK, todoItem)
		return nil
	}
}

func DeleteTodoItem(c *gin.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	todoItem := TodoItem{Id: id}
	
	if err := DBHelper.First(&todoItem).Error; err != nil {
		return UnknownError(err.Error())
	}

	if err := DBHelper.Delete(&todoItem).Error; err != nil {
		return UnknownError(err.Error())
	}else {
		c.JSON(http.StatusOK, todoItem)
		return nil
	}
}