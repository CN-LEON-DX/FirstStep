package ginitem

import (
	"awesomeProject/common"
	"awesomeProject/module/item/biz"
	"awesomeProject/module/item/storage"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetItemById(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.NewCustomError(err, "Invalid ID format", err.Error(), "INVALID_ID"))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewGetItemBiz(store)

		data, err := business.GetItemById(c.Request.Context(), id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				panic(common.NewCustomError(err, "Cannot find item", "Record not found", "ITEM_NOT_FOUND"))
			}
			panic(common.ErrDB(err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
