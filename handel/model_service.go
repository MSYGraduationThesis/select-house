package handel

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"select-house/request"
	"select-house/service"
)

const (
	ResQuest = "request.html"
	Response = "response.html"
)

func GetHouseModelService(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, ResQuest, gin.H{
		"title": "ResQuest",
	})
}

func PostHouseModelService(ctx *gin.Context) {
	req := request.HouseModel{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := service.Service房产投资模型(req.TotalPrice, req.TotalSize, req.ShouFuBili, req.QiShuiLv, req.YinHuaLv, req.Zujin, req.DaiKuanNianXian,
		req.DaiKuanLiLv, req.WuYeFeiDanJia, req.LiCaiLiLv, req.TuiSuanNianFen)

	ctx.HTML(http.StatusOK, Response, gin.H{
		"result":            result,
		"tui_suan_nian_fen": req.TuiSuanNianFen,
	})
}
