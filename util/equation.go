package util

import (
	"fmt"
	"strings"
)

var Operands = map[int]string{
	1: "X",
	2: "Y",
	3: "Z",
}

const (
	SideOne SideCount = iota // 未知数出现在一边
	SideTwo                  // 未知数出现在两边

	BracketNeed      Bracket = iota // 需要括号
	BracketNoNeed                   // 不需要括号
	BracketUncertain                // 不确定(随机决定是否需要)
)

type (
	SideCount int
	Bracket   int
)

type EquationOneOperand struct {
	Side           SideCount // 未知数出现在一边还是两边
	bracket        bool      // 是否需要括号
	Value          int       // 未知数的值
	coefficientOne int       // 第一系数
	coefficiectTwo int       // 第二系数
	sum            int       // 单边的值
	equation1      string    // 公式1
	equation2      string    // 公式2
}

func NewEquationOneOperand(side SideCount, bracket Bracket, minValue, maxValue int, minCoefficient, maxCoefficient int) *EquationOneOperand {
	e := &EquationOneOperand{
		Side:  side,
		Value: RandInt(minValue, maxValue),
	}
	if bracket == BracketUncertain {
		e.bracket = RandBool()
	} else if bracket == BracketNeed {
		e.bracket = true
	} else {
		e.bracket = false
	}
	// 第一章节：生成公式1
	e.equation1, e.sum = NewEquation(e.Side, e.bracket, e.Value, e.Value, minCoefficient, maxCoefficient)

	// 第二章节：反推公式2
	// 第一步：确定单双边
	if side == SideOne {
		// 单边时按几率配置公式2

		// probability := RandInt(1, 100)

		// if probability <= 20 {
		// 	// 几率20%，公式2为公式1的值
		// 	e.equation2 = fmt.Sprintf("%d", e.sum)
		// } else if probability <= 60 {
		// 	// 几率40%，公式2为公式1的值加一个小数
		// }

		e.equation2 = Calculate(e.sum, 20, 40, 40)
	} else {
		// 双边时反向生成公式2
	}

	return e
}

// 构造一个计算等式(不含未知数)
// numberPercent 数字出现的几率
// simplePercent 简单运算出现的几率
// complexPercent 复杂运算出现的几率
func Calculate(sum int, numberPercent, simplePercent, complexPercent int) string {
	probability := RandInt(1, 100)

	if probability <= numberPercent {
		// 数字
		return fmt.Sprintf("%d", sum)
	} else if probability <= numberPercent+simplePercent {
		// 简单运算
		return CalculateSimple(sum)
	} else {
		// 复杂运算

		// 生成括号内的计算式
		subSum := RandIntExclude(2, 10, sum)
		equation := fmt.Sprintf("(%s)", CalculateSimple(subSum))
		// 生成系数
		coefficient := RandIntExclude(-13, 9, -1, 0, 1)
		// 融入系数
		subSum = coefficient * subSum
		equation = fmt.Sprintf("%d%s", coefficient, equation)
		// 判断系数正负
		if coefficient < 0 {
			// 生成减法运算
			equation = fmt.Sprintf("%d %s", -1*subSum+sum, equation)
		} else {
			// 生成加法运算

			// 保证 sum > subSum
			if sum < subSum {
				sum, subSum = subSum, sum
			}

			// 生成加法运算
			if RandBool() {
				// 加到前面
				equation = fmt.Sprintf("%d + %s", sum-subSum, equation)
			} else {
				// 加到后面
				equation = fmt.Sprintf("%s+ %d", equation, sum-subSum)
			}

		}

		return equation
	}
}

// 构造简单计算等式
func CalculateSimple(sum int) string {
	isPlus := RandBool()
	if sum == 1 {
		isPlus = false
	}
	if isPlus {
		// 加法
		value := RandInt(1, sum-1)
		return fmt.Sprintf("%d + %d", value, sum-value)
	} else {
		// 减法
		value := RandInt(sum+1, sum*2)
		return fmt.Sprintf("%d - %d", value, value-sum)
	}
}

// 反向构造一个公式
func ReverseEquation(sum int) string {
	return ""
}

// 构造一个公式
func NewEquation(side SideCount, bracket bool, minValue, maxValue int, minCoefficient, maxCoefficient int) (string, int) {
	var sum int
	var equation string
	value := RandInt(minValue, maxValue)

	// 第一章节：生成公式1
	// 第一步：确定系数
	coefficientOne := RandIntExcludeZero(minCoefficient, maxCoefficient)
	sum = coefficientOne * value
	if coefficientOne != -1 {
		equation = fmt.Sprintf("%d%s", coefficientOne, Operands[1])
	} else {
		equation = fmt.Sprintf("-%s", Operands[1])
	}

	// 第二步：确定前后整数
	if sum < 0 {
		// 如果为负，则需要左置一个大数
		left := RandInt(sum*-1, sum*-10)
		sum = left + sum
		equation = fmt.Sprintf("%d %s", left, equation)
	} else {
		// 如果为正，则需要随机选择前后置一个数
		if RandBool() {
			// 前置一个数
			left := RandInt(2, sum*3)
			sum = left + sum
			equation = fmt.Sprintf("%d + %s", left, equation)
		} else {
			// 后置一个数
			right := RandInt(1, sum-1)
			sum = sum + right
			equation = fmt.Sprintf("%s -%d", equation, right)
		}
	}

	// 第三步：确定括号
	if bracket {
		// 需要增加括号
		equation = fmt.Sprintf("(%s)", equation)

		// 第四步：确定括号系数
		coefficientBracket := RandIntExcludeZero(-9, 9)

		// 融入括号系数
		sum = sum * coefficientBracket
		switch coefficientBracket {
		case -1:
			equation = fmt.Sprintf("-%s", equation)
		case 1:
			equation = fmt.Sprintf("%s", equation)
		default:
			equation = fmt.Sprintf("%d%s", coefficientBracket, equation)
		}

		// 第五步：确定括号外大数
		if sum < 0 {
			// 如果为负，则需要左置一个大数
			left := RandInt(sum*-1, sum*-4)
			sum = left + sum
			equation = fmt.Sprintf("%d %s", left, equation)
		} else {
			// 如果为正，则需要随机选择前后置一个数或不置数
			switch RandInt(-1, 2) {
			case -1:
				// 后置一个减去的小数
				right := RandInt(1, sum-1)
				sum = sum - right
				equation = fmt.Sprintf("%s -%d", equation, right)

			case 1:
				// 前置一个加上的小数
				left := RandInt(2, sum*2)
				sum = left + sum
				equation = fmt.Sprintf("%d + %s", left, equation)
			case 2:
				// 后置一个加上的小数
				right := RandInt(2, sum*2)
				sum = sum + right
				equation = fmt.Sprintf("%s + %d", equation, right)
			default:
				// case 0 或其他数时 不置数
			}
		}

	}
	return equation, sum
}

// 打印公式
func (e *EquationOneOperand) String() string {
	if RandBool() {
		return Standardize(fmt.Sprintf("%s = %s", e.equation1, e.equation2))
	} else {
		return Standardize(fmt.Sprintf("%s = %s", e.equation2, e.equation1))
	}
}

// 规范公式中的空格
func Standardize(equation string) string {
	var buf strings.Builder
	var previousCharIsMinux bool
	// 遍历字符
	for _, c := range equation {
		if previousCharIsMinux && c != ' ' {
			buf.WriteRune(' ')
			previousCharIsMinux = false
		}
		buf.WriteRune(c)
		if c == '-' {
			previousCharIsMinux = true
		}
	}

	return buf.String()
}
