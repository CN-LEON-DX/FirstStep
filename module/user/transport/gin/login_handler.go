package ginuser

import (
	"awesomeProject/common"
	"awesomeProject/component/tokenprovider/jwt"
	"awesomeProject/module/user/biz"
	"awesomeProject/module/user/model"
	"awesomeProject/module/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData model.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(400, common.NewCustomError(
				err,
				err.Error(),
				"Your input is invalid!",
				"ERR_PASS_EMAIL"))
			return
		}

		tokenProvider := jwt.NewTokenJWTProvider("jwt", "sudo_key")

		store := storage.NewSQLStore(db)
		md5 := common.NewMd5Hash()

		bussiness := biz.NewLoginBussiness(store, tokenProvider, md5, 60*60*24*30)

		account, err := bussiness.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			c.JSON(400, common.NewCustomError(
				err,
				err.Error(),
				"Your pass or email wromg !",
				"ERR_PASS_EMAIL"))
			return
		}

		c.JSON(200, common.SimpleSuccessResponse(account))
	}
}
