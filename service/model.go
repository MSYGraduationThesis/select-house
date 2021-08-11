package service

import (
	"fmt"
	"math"
	"select-house/model"
)

// 金钱单位: W  面积单位: m^2

type 房子收益计算 struct {
	房子模型
	买房子投入模型
	理财收益模型
	//计算时长
	推理年份 float64
}

type 房子模型 struct {
	//基础属性
	总价  float64
	总面积 float64
	单价  float64

	契税费率 float64
	印花税  float64

	//每月房贷
	//是否贷款   bool
	首付比例   float64
	首付     float64
	贷款年限   float64
	贷款利率   float64
	贷款月利率  float64
	每月房贷不准 float64
	房贷总额度  float64
	房贷已经支出 float64

	//租金
	租金 float64

	//贷款

	装修费用合计  float64
	物业费用月单价 float64
}

type 理财收益模型 struct {
	理财利率  float64
	理财收益  float64
	理财本息和 float64
}

type 买房子投入模型 struct {
	投入合计 float64
	税费
	装修管理费用
}
type 税费 struct {
	税费总计 float64
	契税   float64
	印花税  float64
}
type 装修管理费用 struct {
	装修管理费用总计 float64
	装修费      float64
	外推N年物业费  float64
}

func (f 房子模型) 房子初始化() 房子模型 {
	f.贷款月利率 = f.贷款利率 / float64(12)
	f.单价 = f.总价 / f.总面积
	f.首付 = f.总价 * f.首付比例
	return f
}

func (m *房子收益计算) 计算贷款() 房子收益计算 {
	//分子 := (房子模型1.总价-  房子模型1.首付) × (1+房子模型1.贷款月利率)^(12*年限) *房子模型1.贷款月利率)
	分子 := (m.房子模型.总价 - m.房子模型.首付) * math.Pow((1+m.房子模型.贷款月利率), float64((12*m.贷款年限))) * m.房子模型.贷款月利率
	分母 := math.Pow(1+m.房子模型.贷款月利率, (12*m.贷款年限)) - 1
	每月房贷 := 分子 / 分母
	m.房子模型.每月房贷不准 = 每月房贷
	m.房子模型.房贷已经支出 = 每月房贷 * m.推理年份 * 12
	return *m
}

func (m 房子收益计算) 计算装修管理() 房子收益计算 {
	m.装修管理费用.装修费 = m.房子模型.装修费用合计
	m.装修管理费用.外推N年物业费 = m.房子模型.物业费用月单价 * 12 * m.房子模型.总面积 * m.推理年份
	m.装修管理费用总计 = m.装修管理费用.装修费 + m.装修管理费用.外推N年物业费
	return m
}

func (m 房子收益计算) 计算税费() 房子收益计算 {
	m.税费.契税 = m.房子模型.契税费率 * m.房子模型.总价
	m.税费.印花税 = m.房子模型.印花税 * m.房子模型.总价
	m.税费总计 = m.税费.契税 + m.税费.印花税
	return m
}

func (m 房子收益计算) 计算投入成本合计() 房子收益计算 {
	//投入费用= 房屋首付+其他费用+贷款已经出去
	m.投入合计 = m.房子模型.首付 + m.税费总计 + m.装修管理费用总计 + m.房贷已经支出
	//
	租金 := m.房子模型.租金 * 12 * m.推理年份
	m.投入合计 = m.投入合计 - 租金
	return m
}

func (m 房子收益计算) 计算理财收入() 房子收益计算 {
	m.理财收益 = (m.房子模型.首付 + m.税费总计 + m.装修管理费用总计) * m.理财收益模型.理财利率 * m.推理年份
	m.理财本息和 = (m.房子模型.首付 + m.税费总计 + m.装修管理费用总计) * (m.理财收益模型.理财利率*m.推理年份 + 1)
	return m
}

func New房子模型(总价, 总面积, 首付比例, 契税费率, 印花税, 租金, 贷款年限, 贷款利率, 物业费用月单价 float64) 房子模型 {
	房子 := 房子模型{
		总价:   总价,
		总面积:  总面积,
		首付比例: 首付比例,
		契税费率: 契税费率,
		印花税:  印花税,
		//是否贷款:   true,
		租金:      租金,
		贷款年限:    贷款年限,
		贷款利率:    贷款利率,
		物业费用月单价: 物业费用月单价,
	}
	房子 = 房子.房子初始化()
	return 房子
}
func New合肥房子模型() 房子模型 {
	房子 := 房子模型{
		总价:   2000000,
		总面积:  100,
		单价:   20000,
		契税费率: 0.015,
		印花税:  0.0005,
		首付比例: 0.3,
		//是否贷款:   true,
		租金:      2000,
		贷款年限:    30,
		贷款利率:    0.0588,
		物业费用月单价: 1.26,
	}
	房子 = 房子.房子初始化()
	return 房子
}

func Service房产投资模型(总价, 总面积, 首付比例, 契税费率, 印花税, 租金, 贷款年限, 贷款利率, 物业费用月单价, 理财利率, 推理年限 float64) model.Resp {

	//模型 := New房子模型(总价, 总面积, 首付比例, 契税费率, 印花税, 租金, 贷款年限, 贷款利率, 物业费用月单价)
	模型 := New房子模型(670000, 97, 0.3, 0.015,
		0.0005, 2000, 30, 0.0588, 2)
	money := 房子收益计算{
		房子模型:   模型,
		理财收益模型: 理财收益模型{理财利率: 理财利率},
		推理年份:   推理年限,
	}

	fmt.Println("-------五年为期限测算---")
	fmt.Printf("模型总价:%.4f \n", 模型.总价)
	fmt.Println("模型单价:", 模型.单价)
	fmt.Println("模型首付比例:", 模型.首付比例)
	fmt.Println("模型贷款年限:", 模型.贷款年限)
	money = money.计算贷款()
	fmt.Println("计算贷款支出总额:", money.房贷已经支出)
	money = money.计算税费()
	fmt.Println("计算税费:", money.税费总计)
	money = money.计算装修管理()
	fmt.Println("计算装修管理:", money.装修管理费用总计)
	money = money.计算投入成本合计()
	fmt.Println(fmt.Sprintf("投入成本合计:%.5f ", money.投入合计))
	money = money.计算理财收入()
	fmt.Println(fmt.Sprintf("理财本息和:%.5f ", money.理财本息和))
	// 房子需要x涨幅才能弥补理财收益
	result := model.Resp{}
	X := 0.01
	for {
		X += 0.01
		房子的当前总价 := money.房子模型.单价 * (1 + X) * money.房子模型.总面积
		房子价值增长了 := 房子的当前总价 - money.房子模型.总价
		剩余贷款 := money.房子模型.每月房贷不准 * ((money.贷款年限 - money.推理年份) * 12)
		纯投房子收益 := 房子的当前总价 - 剩余贷款 - money.投入合计
		收益差 := 纯投房子收益 - money.理财收益
		if 收益差 > 0 {
			fmt.Println("-------按照 房子价值增长大于理财收入，为测算标准，结果如下-----")
			fmt.Printf("房价需要增长百分数%.4f：房子的单价达到=%.4f \n", X, money.房子模型.单价*(1+X))
			fmt.Println(fmt.Sprintf("房子的当前总价:%4.f", 房子的当前总价))
			fmt.Println(fmt.Sprintf("房子的投入合计:%4.f", money.投入合计))
			fmt.Println(fmt.Sprintf("剩余贷款:%4.f", 剩余贷款))
			fmt.Println(fmt.Sprintf("房子的总价值增长了:%4.f", 房子价值增长了))
			fmt.Printf("纯投房子收益:%.4f\n", 纯投房子收益)
			fmt.Println("到期理财本息共计:", money.理财本息和)
			fmt.Println("理财收益增长了:", money.理财收益)
			fmt.Println("买房子多收益:", 收益差)
			result = model.Resp{
				ShuiFeiHeJi:          fmt.Sprintf("%.2f", money.税费总计),
				ZhuangXiuWuYeFeiHeji: fmt.Sprintf("%.2f", money.装修管理费用总计),
				HasPaiDaiKuan:        fmt.Sprintf("%.2f", money.房贷已经支出),
				HasPaidZongChengBen:  fmt.Sprintf("%.2f", money.投入合计),
				LiCaiBenXiHe:         fmt.Sprintf("%.2f", money.理财本息和),

				FangJiaXuYaoZengZhangBaiFenShu: fmt.Sprintf("%.3f", X),
				FangJiZongJiaDaDao:             fmt.Sprintf("%.2f", 房子的当前总价),

				ZongJiaZhiZengZhangLe: fmt.Sprintf("%.2f", 房子价值增长了),
				ShengYuDaiKuan:        fmt.Sprintf("%.2f", 剩余贷款),
				ChunTouFangZiShouYi:   fmt.Sprintf("%.2f", 纯投房子收益),
				ChunLiCaiShouYi:       fmt.Sprintf("%.2f", money.理财收益),
				ShouYiCha:             fmt.Sprintf("%.2f", 收益差),
			}

			break
		}
	}
	fmt.Printf("结论：房价需要增长百分数%.4f ：房子的单价达到=%.4f \n", X, money.房子模型.单价*(1+X))
	return result
}
