package api

import (
	"NectarPin/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSystemInfo(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": "200",
			"data": gin.H{
				"NectarpinVersion": constant.NectarpinVersion,
			},
			"msg": "[NectarPin]: 查询成功！",
		},
	)
}
