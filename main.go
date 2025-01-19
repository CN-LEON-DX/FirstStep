package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id"` // tag name json, json encode
	Title       string     `json:"string" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
	//UpdatedAt *time.Time `json:"updated_at"`
}
type TodoItemCreation struct {
	Id          int    `json:"id" gorm:"column:id"` // tag name json, json encode
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}
type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string         { return "todo_items" }
func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }
func (TodoItemUpdate) TableName() string   { return TodoItem{}.TableName() }

func main() {

	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database", db)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("", ListItem(db))
			items.GET("/:id", GetItemById(db))
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

func CreateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var itemData TodoItemCreation
		if err := c.ShouldBind(&itemData); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		if err := db.Create(&itemData).Error; err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": itemData.Id,
		})
	}
}

func GetItemById(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var itemData TodoItem
		// id is value pass by param and err is error thrown if convertion error
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		log.Println(id)
		fmt.Println(id)

		if err := db.Where("id = ?", id).First(&itemData).Error; err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": itemData,
		})
	}
}

func UpdateInfoItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var updateData TodoItemUpdate
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
		c.JSON(http.StatusOK, gin.H{
			"data_return": true,
		})
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
		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(&TodoItemUpdate{Status: &deletedStatus}).Error; err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data_return": true,
		})
	}
}

func ListItem(db *gorm.DB) func(c *gin.Context) {
	log.Println("Run")
	return func(c *gin.Context) {

		var result []TodoItem
		if err := db.Table(TodoItem{}.TableName()).Order("status asc").Find(&result).Error; err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data_return": result,
		})
	}
}
