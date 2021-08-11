package request

type HouseModel struct {
	TotalPrice float64 `json:"total_price"`
	TotalSize  float64 `json:"total_size"`
	//UnitPrice       float64 `json:"unit_price"`
	QiShuiLv        float64 `json:"qi_shui_rate"`
	YinHuaLv        float64 `json:"yin_hua_lv"`
	ShouFuBili      float64 `json:"shou_fu_bili"`
	Zujin           float64 `json:"zujin"`
	DaiKuanNianXian float64 `json:"dai_kuan_nian_xian"`
	DaiKuanLiLv     float64 `json:"dai_kuan_li_lv"`
	WuYeFeiDanJia   float64 `json:"wu_ye_fei_dan_jia"`

	// 理财部分
	LiCaiLiLv float64 `json:"li_cai_li_lv"`
	//
	TuiSuanNianFen float64 `json:"tui_suan_nian_fen"`
}
