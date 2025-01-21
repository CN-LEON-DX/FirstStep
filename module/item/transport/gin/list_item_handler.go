package ginitem

import (
	"awesomeProject/common"
	"awesomeProject/module/item/biz"
	"awesomeProject/module/item/model"
	"awesomeProject/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			model.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		queryString.Paging.Process()

		store := storage.NewSQLStore(db)
		business := biz.NewListItem(store)

		listItem, err := business.ListItem(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(listItem, queryString.Filter, queryString.Paging))
	}
}
