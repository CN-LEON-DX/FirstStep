package ginitem

import (
	"awesomeProject/common"
	"awesomeProject/module/item/biz"
	"awesomeProject/module/item/model"
	"awesomeProject/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func UpdateInfoItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var updateData model.TodoItemUpdate
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}

		if err := c.ShouldBind(&updateData); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(store)

		if err := business.UpdateItemById(c.Request.Context(), id, &updateData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
