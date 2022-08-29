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

	NeedYes       Need = iota // 需要
	NeedNo                    // 不需要
	NeedUncertain             // 不确定(随机决定是否需要)
)

type (
	// 未知数出现在一边还是两边
	SideCount int
	// 需求情况
	Need int
)

// 一元一次方程实体
type EquationOneOperand struct {
	Side           SideCount // 未知数出现在一边还是两边
	Value          int       // 未知数的值(用于公布答案)
	needBracket    bool      // 是否需要括号
	needCefficient bool      // 是否需要给未知数添加系数
	sum            int       // 单边的计算结果值
	equation1      string    // 生成的公式1
	equation2      string    // 生成的公式2
}

// 构造一元一次等式
func NewEquationOneOperand(side SideCount, bracket Need, minValue, maxValue int, minCoefficient, maxCoefficient int) *EquationOneOperand {
	// 单边未知数系数‘
	var coefficient int

	e := &EquationOneOperand{
		Side:  side,
		Value: RandInt(minValue, maxValue),
	}
	if bracket == NeedUncertain {
		e.needBracket = RandBool()
	} else if bracket == NeedYes {
		e.needBracket = true
	} else {
		e.needBracket = false
	}
	// 第一章节：生成公式1
	e.equation1, e.sum, coefficient = NewEquation(Operands[1], e.needBracket, e.Value, minCoefficient, maxCoefficient)

	// 第二章节：反推公式2
	// 第一步：确定单双边
	if side == SideOne {
		// 单边时按几率配置公式2
		e.equation2 = NewCalculate(e.sum, 20, 40, 40)
	} else {
		// 双边时反向生成公式2

		needBracketForReverse := e.needBracket
		if needBracketForReverse {
			// equation1 已经有括号了，equation2 的括号随机决定
			needBracketForReverse = RandBool()
		}
		e.equation2 = ReverseEquation(Operands[1], coefficient, e.Value, e.sum, needBracketForReverse, minCoefficient, maxCoefficient)
	}

	return e
}

// 构造一个数字计算等式(不含未知数)
// numberPercent 数字出现的几率
// simplePercent 简单运算出现的几率
// complexPercent 复杂运算出现的几率
func NewCalculate(sum int, numberPercent, simplePercent, complexPercent int) string {
	probability := RandInt(1, 100)

	if probability <= numberPercent {
		// 数字
		return fmt.Sprintf("%d", sum)
	} else if probability <= numberPercent+simplePercent {
		// 简单运算
		return NewCalculateSimple(sum)
	} else {
		// 复杂运算

		// 生成括号内的计算式
		subSum := RandIntExclude(2, 19, sum)
		equation := fmt.Sprintf("(%s)", NewCalculateSimple(subSum))
		// 生成系数
		coefficient := RandIntExclude(-13, 9, 0)
		// 融入系数
		subSum = coefficient * subSum
		switch coefficient {
		case -1:
			equation = fmt.Sprintf("-%s", equation)
		case 1:
			equation = fmt.Sprintf("%s", equation)
		default:
			equation = fmt.Sprintf("%d%s", coefficient, equation)
		}

		// 判断系数正负
		if coefficient < 0 {
			// 生成减法运算
			equation = fmt.Sprintf("%d %s", -1*subSum+sum, equation)
		} else {
			// 根据 sum 和 subSum 的大小决定加减法运算符

			if sum > subSum {
				// 生成加法运算
				if RandBool() {
					// 加到前面
					equation = fmt.Sprintf("%d + %s", sum-subSum, equation)
				} else {
					// 加到后面
					equation = fmt.Sprintf("%s + %d", equation, sum-subSum)
				}
			} else {
				// 生成减法运算
				equation = fmt.Sprintf("%s - %d", equation, subSum-sum)
			}
		}

		return equation
	}
}

// 构造一个数字计算等式(包含运算符，但不含未知数)
func NewCalculateSimple(sum int) string {
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

// 构造一元一次公式
// 返回值：公式，公式计算结果值，未知数的系数
func NewEquation(operand string, needBracket bool, value int, minCoefficient, maxCoefficient int) (string, int, int) {
	var sum int
	var equation string
	coefficientBracket := 1 // 括号的系数

	// 第一章节：生成公式1
	// 第一步：确定未知数的
	coefficient := RandIntExcludeZero(minCoefficient, maxCoefficient)
	sum = coefficient * value
	switch coefficient {
	case 1:
		equation = operand
	case -1:
		equation = fmt.Sprintf("-%s", operand)
	default:
		equation = fmt.Sprintf("%d%s", coefficient, operand)
	}

	// 第二步：确定前后整数
	if sum < 0 {
		// 如果为负，则需要左置一个数
		left := RandIntExclude(sum*-1, sum*-1+10, sum)
		sum = left + sum
		equation = fmt.Sprintf("%d %s", left, equation)
	} else {
		// 如果为正，则需要随机选择前后置一个数
		add := RandInt(2, 17)
		sum = add + sum
		if RandBool() {
			// 前置一个加数
			equation = fmt.Sprintf("%d + %s", add, equation)
		} else {
			// 后置一个加数
			equation = fmt.Sprintf("%s + %d", equation, add)
		}
	}

	// 第三步：确定括号
	if needBracket {
		// 需要增加括号
		equation = fmt.Sprintf("(%s)", equation)

		// 第四步：确定括号系数
		coefficientBracket = RandIntExcludeZero(-9, 9)

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
	return equation, sum, coefficient * coefficientBracket
}

// 打印一元一次等式
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
	var previousCharIsSpace bool
	// 遍历字符
	for _, c := range equation {
		if previousCharIsSpace && c == ' ' {
			// 去掉连续的空格
			continue
		}
		if previousCharIsMinux && c != ' ' {
			// 在减号后面加一个空格
			buf.WriteRune(' ')
			previousCharIsMinux = false
		}
		buf.WriteRune(c)
		if c == '-' {
			previousCharIsMinux = true
		}
		if c == ' ' {
			previousCharIsSpace = true
		} else {
			previousCharIsSpace = false
		}
	}

	return buf.String()
}

// 反向构造一元一次公式
//
// 参数：
//    operand: 未知数符号
//    excludeCoefficient: 未知数的系数，反推公式中系数应该避免这个值
//    value: 未知数的值
//    sum: 公式的值
//    needBracket: 是否需要括号
//    minCoefficient: 系数范围最小值
//    maxCoefficient: 系数范围最大值
// 返回：
//    公式字符串
func ReverseEquation(operand string, excludeCoefficient, value, sum int, needBracket bool, minCoefficient, maxCoefficient int) string {
	var coefficientNew, coefficientBracket int
	var sumNew int
	var equation string

	// 第一步：确定未知数系数
	coefficientNew = RandIntExclude(minCoefficient, maxCoefficient, 0, excludeCoefficient)
	coefficientBracket = RandIntExclude(minCoefficient, maxCoefficient, 0, excludeCoefficient)
	for coefficientNew*coefficientBracket == excludeCoefficient {
		coefficientBracket = RandIntExclude(minCoefficient, maxCoefficient, 0, excludeCoefficient)
	}

	sumNew = coefficientNew * value
	switch coefficientNew {
	case 1:
		equation = operand
	case -1:
		equation = fmt.Sprintf("-%s", operand)
	default:
		equation = fmt.Sprintf("%d%s", coefficientNew, operand)
	}

	// 第二步：确定前后整数
	if sumNew < 0 {
		// 如果为负，则需要左置一个大数
		left := RandInt(sumNew*-1, sumNew*-1+10)
		sumNew = left + sumNew
		equation = fmt.Sprintf("%d %s", left, equation)
	} else {
		// 如果为正，则需要随机选择前后置一个数
		add := RandInt(2, 17)
		sumNew = add + sumNew
		if RandBool() {
			// 前置一个加数
			equation = fmt.Sprintf("%d + %s", add, equation)
		} else {
			// 后置一个加数
			equation = fmt.Sprintf("%s + %d", equation, add)
		}
	}

	// 第三步：确定括号
	if needBracket {
		// 需要增加括号
		equation = fmt.Sprintf("(%s)", equation)

		// 融入括号系数
		sumNew = sumNew * coefficientBracket
		switch coefficientBracket {
		case -1:
			equation = fmt.Sprintf("-%s", equation)
		case 1:
			equation = fmt.Sprintf("%s", equation)
		default:
			equation = fmt.Sprintf("%d%s", coefficientBracket, equation)
		}

		// 第五步：确定括号外大数
		if sumNew < 0 {
			// 如果为负，则需要左置一个大数
			equation = fmt.Sprintf("%d %s", -1*sumNew+sum, equation)
		} else {
			// 如果为正，需要判断 sum 和 sumNew 的大小关系
			if sum < sumNew {
				// 如果 sum 小于 sumNew，则需要右减两数的差
				equation = fmt.Sprintf("%s - %d", equation, sumNew-sum)
			} else {
				// 如果 sum 大于 sumNew，则需要加上两数的差
				if RandBool() {
					// 前置一个加数
					equation = fmt.Sprintf("%d + %s", sum-sumNew, equation)
				} else {
					// 后置一个加数
					equation = fmt.Sprintf("%s + %d", equation, sum-sumNew)
				}
			}
		}
	} else {
		// 不需要增加括号，则需要判断 sum 和 sumNew 的大小关系
		if sum < sumNew {
			// 如果 sum 小于 sumNew，则需要右减两数的差
			equation = fmt.Sprintf("%s - %d", equation, sumNew-sum)
		} else {
			// 如果 sum 大于 sumNew，则需要加上两数的差
			if RandBool() {
				// 前置一个加数
				equation = fmt.Sprintf("%d + %s", sum-sumNew, equation)
			} else {
				// 后置一个加数
				equation = fmt.Sprintf("%s + %d", equation, sum-sumNew)
			}
		}
	}

	return equation
}
