package main

import (
	"fmt"

	"github.com/woshizilong/equation/util"
)

func main() {
	test4()
}

func test0() {
	for i := 1; i <= 10; i++ {
		eq := util.NewEquationOneOperand(util.SideTwo, util.NeedYes, 2, 5, -2, 2)
		fmt.Printf("%d  %s\n", eq.Value, eq.String())
	}
}

func test1() {
	ep := util.NewExamPaper_1U1P2S("一元一次方程", "一元一次方程-0829")
	pdf, pdfWithAnswer := ep.Generate(util.AnswerPageSeparate)
	fmt.Printf("生成了试卷\n  %s\n  %s\n", pdf, pdfWithAnswer)
}

func test2() {
	for z := 1; z <= 3; z++ {
		for x := 0; x <= 3; x++ {
			n := util.NewDigital(z, x)
			fmt.Printf("整数%d位，小数%d位 -- %s\n", z, x, n.ValueString)
		}
		fmt.Println("====================================")
	}
}

func test3() {
	for z := 1; z <= 3; z++ {
		for x := 0; x <= 3; x++ {
			n := util.NewMultiplyingDecimals(z, z, x, x)
			fmt.Printf("%s = %s\n", n.String, n.Product.String())
		}
		fmt.Println("====================================")
	}
}

func test4() {
	ep := util.NewDecimalMultiplicationExamPaper_2D("小数乘法练习", "两位小数-0905")
	pdf, pdfWithAnswer := ep.Generate(util.AnswerPageSeparate)
	fmt.Printf("生成了试卷\n  %s\n  %s\n", pdf, pdfWithAnswer)
}
