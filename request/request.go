package request

type HouseModel struct {
	TotalPrice float64 `form:"total_price"`
	TotalSize  float64 `form:"total_size"`
	//UnitPrice       float64 `json:"unit_price"`
	QiShuiLv        float64 `form:"qi_shui_rate"`
	YinHuaLv        float64 `form:"yin_hua_lv"`
	ShouFuBili      float64 `form:"shou_fu_bili"`
	Zujin           float64 `form:"zujin"`
	DaiKuanNianXian float64 `form:"dai_kuan_nian_xian"`
	DaiKuanLiLv     float64 `form:"dai_kuan_li_lv"`
	WuYeFeiDanJia   float64 `form:"wu_ye_fei_dan_jia"`

	// 理财部分
	LiCaiLiLv float64 `form:"li_cai_li_lv"`
	//
	TuiSuanNianFen float64 `form:"tui_suan_nian_fen"`
}
