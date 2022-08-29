package util

import "fmt"

const (
	// 考试难度
	ExamLevelEasy   ExamLevel = iota // 难度：简单
	ExamLevelMedium                  // 难度：中等
	ExamLevelHard                    // 难度：困难

	// 是否需要答案页
	AnswerPageInclude  AnswerPage = iota // 需要答案页
	AnswerPageExclude                    // 不需要答案页
	AnswerPageSeparate                   // 分开生成答案页和试卷页
)

type (
	ExamLevel  int
	AnswerPage int
)

type ExamPaper struct {
	Title          string    // 考试题目
	Level          ExamLevel // 考试难度
	Number         int       // 考试题目数量
	PerPage        int       // 每行题目数量
	Power          int       // 几次方
	Unknowns       int       // 未知数个数
	Side           SideCount // 双边还是单边
	ValueMin       int       // 结果最小值
	ValueMax       int       // 结果最大值
	CoefficientMin int       // 系数最小值
	CoefficientMax int       // 系数最大值
	Filename       string    // 考试题目的pdf文件路径
}

// 生成试卷
func (ep *ExamPaper) Generate(needAnswer AnswerPage) (string, string) {
	// 生成考试题目
	var equations [][]string
	var answers string

	var lineEquations []string

	for i := 1; i <= ep.Number; i++ {
		e := NewEquationOneOperand(ep.Side, NeedUncertain, ep.ValueMin, ep.ValueMax, ep.CoefficientMin, ep.CoefficientMax)
		lineEquations = append(lineEquations, fmt.Sprintf("(%d). %s", i, e.String()))

		answer := fmt.Sprintf("(%d). %d", i, e.Value)
		answers = fmt.Sprintf("%s %-10s", answers, answer)

		if i%ep.PerPage == 0 {
			line := make([]string, ep.PerPage)
			copy(line, lineEquations)
			equations = append(equations, line)
			// 清空
			lineEquations = lineEquations[:0]
		}
	}
	// 判断是否有剩余的题目
	if len(lineEquations) > 0 {
		for i := len(lineEquations); i < ep.PerPage; i++ {
			lineEquations = append(lineEquations, "")
		}
		equations = append(equations, lineEquations)
	}
	// 生成考试题目的pdf文件
	var outputFilename string
	var outputFilenameWithAnswer string
	switch needAnswer {
	case AnswerPageInclude:
		outputFilenameWithAnswer = fmt.Sprintf("pdfs/%s-含答案.pdf", ep.Filename)
		OutputPdfWithAnswer(outputFilenameWithAnswer, ep.Title, equations, answers)
	case AnswerPageExclude:
		outputFilename = fmt.Sprintf("pdfs/%s.pdf", ep.Filename)
		OutputPdf(outputFilename, ep.Title, equations)
	case AnswerPageSeparate:
		outputFilename = fmt.Sprintf("pdfs/%s.pdf", ep.Filename)
		OutputPdf(outputFilename, ep.Title, equations)
		outputFilenameWithAnswer = fmt.Sprintf("pdfs/%s-含答案.pdf", ep.Filename)
		OutputPdfWithAnswer(outputFilenameWithAnswer, ep.Title, equations, answers)
	}

	return outputFilename, outputFilenameWithAnswer
}

// 一元一次方程:生成一个单边的一次方的试卷
func NewExamPaper_1U1P1S(title, filename string) *ExamPaper {
	return &ExamPaper{
		Title:          title,
		Level:          ExamLevelEasy,
		Number:         16,
		PerPage:        2,
		Power:          1,
		Unknowns:       1,
		Side:           SideOne,
		ValueMin:       2,
		ValueMax:       5,
		CoefficientMin: -9,
		CoefficientMax: 9,
		Filename:       filename,
	}
}

// 一元一次方程:生成一个双边的一次方的试卷
func NewExamPaper_1U1P2S(title, filename string) *ExamPaper {
	return &ExamPaper{
		Title:          title,
		Level:          ExamLevelEasy,
		Number:         16,
		PerPage:        2,
		Power:          1,
		Unknowns:       1,
		Side:           SideTwo,
		ValueMin:       2,
		ValueMax:       9,
		CoefficientMin: -4,
		CoefficientMax: 7,
		Filename:       filename,
	}
}
