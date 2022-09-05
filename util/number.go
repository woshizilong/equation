package util

import (
	"github.com/shopspring/decimal"
)

type Digital struct {
	Integer     int             // 整数位数
	Decimal     int             // 小数位数
	Value       decimal.Decimal // 计算用值
	ValueString string          // 显示用值
}

// 随机生成一个指定位数的高精度小数
func NewDigital(i, d int) Digital {
	var zsw int // 整数位数
	var xsw int // 小数位数

	digital := &Digital{}

	// 确定整数位和小数位
	if i <= 1 {
		zsw = 1
	} else {
		zsw = i
	}
	if d < 0 {
		xsw = 0
	} else {
		xsw = d
	}
	digital.Integer = zsw
	digital.Decimal = xsw

	// 生成整数部分
	if zsw == 1 {
		// 生成范围 0-9
		digital.Value = decimal.NewFromInt(int64(RandInt(0, 9)))
	} else {
		// 生成范围 1-9
		digital.Value = decimal.NewFromInt(int64(RandInt(1, 9)))
	}

	// 补足整数位数
	for i := 1; i < zsw; i++ {
		digital.Value = digital.Value.Mul(decimal.NewFromInt(10)).Add(decimal.NewFromInt(int64(RandInt(0, 9))))
	}

	// 判断是否需要生成小数部分
	if xsw > 0 {
		div := decimal.NewFromInt(10)
		for i := 1; i <= xsw; i++ {
			if i == xsw {
				// 最后一位小数
				digital.Value = digital.Value.Add(decimal.NewFromInt(int64(RandInt(1, 9))).Div(div))
			} else {
				// 非最后一位小数
				digital.Value = digital.Value.Add(decimal.NewFromInt(int64(RandInt(0, 9))).Div(div))
			}
			div = div.Mul(decimal.NewFromInt(10))
		}
	}

	digital.ValueString = digital.Value.String()

	return *digital
}
