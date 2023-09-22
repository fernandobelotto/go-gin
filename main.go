package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	Name  string
	Value uint
}

type EntityBody struct {
	Name  string `json:"name"`
	Value uint   `json:"value"`
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		AllowGlobalUpdate: true,
	})

	if err != nil {
		fmt.Println("error while opening db")
		return
	}

	db.AutoMigrate(&Entity{})

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {

		var entityBody EntityBody

		if errA := c.ShouldBind(&entityBody); errA != nil {
			c.JSON(http.StatusOK, `the body should be name and value`)
			return
		}

		db.Create(&Entity{
			Name:  entityBody.Name,
			Value: entityBody.Value,
		})

		c.JSON(200, entityBody)
	})

	r.GET("/", func(c *gin.Context) {
		var entities []Entity

		db.Find(&entities)

		c.JSON(200, entities)
	})

	r.GET("/:id", func(c *gin.Context) {

		id := c.Param("id")

		var entity Entity

		db.Find(&entity, id)

		c.JSON(200, entity)
	})

	r.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")

		var entity Entity

		entityBody := EntityBody{}

		db.First(&entity, id)

		if errA := c.ShouldBind(&entityBody); errA != nil {
			c.JSON(http.StatusOK, `the body should be name and value`)
		}

		db.Model(&entity).Updates(entityBody)

		c.JSON(200, entity)
	})

	r.DELETE("/:id", func(c *gin.Context) {

		db.Delete(&Entity{})

		c.Status(200)
	})

	r.DELETE("/", func(c *gin.Context) {

		var entities []Entity

		db.Delete(&entities)

		c.JSON(200, entities)
	})

	r.Run()
}
