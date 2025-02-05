package gin

import (
	"awesomeProject/common"
	biz "awesomeProject/module/user/biz"
	"awesomeProject/module/user/model"
	"awesomeProject/module/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterHandler(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		//db := sc.MustGet(common.DBMain).(*gorm.DB)
		var data model.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		md5 := common.NewMd5Hash()

		biz := biz.NewRegisterBussiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		//data.Mask(common.DbTypeUser)

		c.JSON(
			http.StatusOK, common.SimpleSuccessResponse(data.Id))

	}
}
