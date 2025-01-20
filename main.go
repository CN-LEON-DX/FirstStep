package main

import (
	"awesomeProject/common"
	"awesomeProject/module/item/model"
	ginitem "awesomeProject/module/item/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// debug
	db = db.Debug()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database", db)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ListItem(db))
			items.GET("/:id", ginitem.GetItemById(db))
			items.PUT("")
			items.PATCH("/:id", UpdateInfoItem(db))
			items.DELETE("/:id", DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3009") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func UpdateInfoItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var updateData model.TodoItemUpdate
		// id is value pass by param and err is error thrown if convertion error
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		log.Println(id)

		if err := c.ShouldBind(&updateData); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).Updates(&updateData).Error; err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}

func DeleteItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		log.Println(id)

		deletedStatus := "Deleted"

		// if we dont pass the item the db won't recognize the table of DB
		if err := db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).Updates(&model.TodoItemUpdate{Status: &deletedStatus}).Error; err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}

func ListItem(db *gorm.DB) func(c *gin.Context) {
	log.Println("Run")
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Process() // change current paging when query

		var result []model.TodoItem

		if err := db.Table(model.TodoItem{}.TableName()).
			Count(&paging.Total).
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Order("status asc").Find(&result).Error; err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
