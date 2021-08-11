package model

type Resp struct {
	ShuiFeiHeJi                    string `json:"shui_fei"`
	税费合计                           string
	ZhuangXiuWuYeFeiHeji           string `json:"zhuangxiu_wu_ye_fei"`
	装修物业费合计                        string
	HasPaiDaiKuan                  string `json:"had_pai_dai_kuan"`
	已经支付贷款                         string
	HasPaidZongChengBen            string //房屋首付+其他费用+贷款已经出去
	已经支付总成本                        string
	LiCaiBenXiHe                   string
	理财本息和                          string
	FangJiaXuYaoZengZhangBaiFenShu string
	房价需要增长百分数                      string
	FangJiZongJiaDaDao             string
	房子总价达到                         string
	ZongJiaZhiZengZhangLe          string
	房子的总价值增长了                      string
	ShengYuDaiKuan                 string
	剩余贷款                           string
	ChunTouFangZiShouYi            string
	纯投房子收益                         string
	ChunLiCaiShouYi                string
	纯理财收益                          string
	ShouYiCha                      string
	收益差                            string
}
