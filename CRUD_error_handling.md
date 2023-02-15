config.go
```go
package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
    var err error
    DB, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
    if err != nil {
        return err
    }
    return nil
}

```

```go
    if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Todo"})
        }
        return
    }


    if err := config.DB.Find(&todos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Todos"})
        return
    }


    if err := config.DB.Create(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating Todo"})
        return
    }


    if err := config.DB.Model(&todo).Updates(models.Todo{Title: input.Title, Description: input.Description}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating Todo"})
        return
    }

    if err := config.DB.Delete(&todo, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting Todo"})
        return
    }

```


```go

    result := config.DB.Create(&todo)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }


    result := config.DB.Model(&todo).Updates(models.Todo{Title: input.Title, Description: input.Description})
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    result := config.DB.Delete(&todo, c.Param("id"))
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }


    result := config.DB.Find(&todos)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    result := config.DB.First(&todo, c.Param("id"))
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        }
        return
    }

```
