package main

import (
	"fmt"

	"github.com/woshizilong/equation/util"
)

func main() {
	// for i := 0; i < 20; i++ {
	// 	value := util.RandIntExcludeZero(0, 0)
	// 	fmt.Println(value)
	// }

	var equations [][]string
	var answers string

	// 每行出几道题
	perPage := 2

	var lineEquations []string

	for i := 1; i <= 16; i++ {
		// fmt.Printf("%d. 答案：%d\t\t公式：%s\n", i+1, e.Value, e.String())
		e := util.NewEquationOneOperand(util.SideOne, util.NeedUncertain, 2, 22, -9, 9)
		lineEquations = append(lineEquations, fmt.Sprintf("(%d). %s", i, e.String()))

		answer := fmt.Sprintf("(%d). %d", i, e.Value)
		answers = fmt.Sprintf("%s %-10s", answers, answer)

		if i%perPage == 0 {
			line := make([]string, perPage)
			copy(line, lineEquations)
			equations = append(equations, line)
			// 清空
			lineEquations = lineEquations[:0]
		}
	}
	// 判断是否有剩余的题目
	if len(lineEquations) > 0 {
		for i := len(lineEquations); i < perPage; i++ {
			lineEquations = append(lineEquations, "")
		}
		equations = append(equations, lineEquations)
	}

	util.OutputPdf("pdfs/T1.pdf", "一元一次方程式练习题", equations)
	util.OutputPdfWithAnswer("pdfs/T1-含答案.pdf", "一元一次方程式练习题", equations, answers)
}
