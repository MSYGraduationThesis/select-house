package handel

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ResQuest = "request.html"
	Response = "response.html"
)

func GetHouseModelService(ctx *gin.Context) {
	//ctx.HTML(http.StatusOK, "request.html",ResQuest)
	ctx.HTML(http.StatusOK, ResQuest, gin.H{
		"title": "ResQuest",
	})
}

func PostHouseModelService(ctx *gin.Context) {
	//req := request.HouseModel{}
	//if err := ctx.ShouldBindJSON(&req); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}

	//总价, 总面积, 首付比例, 契税费率, 印花税, 租金, 贷款年限, 贷款利率, 物业费用月单价 ,理财利率,推理年限
	//
	//result := service.Service房产投资模型(req.TotalPrice,req.TotalSize,req.ShouFuBili,req.QiShuiLv,req.YinHuaLv,req.Zujin,req.DaiKuanNianXian,
	//	req.DaiKuanLiLv,req.WuYeFeiDanJia, req.LiCaiLiLv, req.TuiSuanNianFen)
	//

	ctx.String(http.StatusOK, "智能推理结果%v", nil)
	//ctx.HTML(http.StatusOK, "智能推理结果", nil)

}
