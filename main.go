package main

import (
	"fmt"

	"github.com/woshizilong/equation/util"
)

func main() {
	test1()
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
