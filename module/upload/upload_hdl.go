package upload

import (
	"awesomeProject/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Upload(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			c.JSON(400, err)
			return
		}

		dst := fmt.Sprintf("static/%d%s", time.Now().UTC().UnixNano(), fileHeader.Filename)
		if err := c.SaveUploadedFile(
			fileHeader,
			dst,
		); err != nil {
			c.JSON(400, err)
			return
		}

		img := common.Image{
			Id:        0,
			Url:       dst,
			Width:     100,
			Height:    100,
			CloudName: "local",
			Extension: "",
		}

		img.Fulfill("http://localhost:3009")

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
