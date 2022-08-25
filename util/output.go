package util

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func OutputPdf(filePath string, title string, items [][]string) {
	outputPdf(filePath, title, items, "", false)
}

func OutputPdfWithAnswer(filePath string, title string, items [][]string, answers string) {
	outputPdf(filePath, title, items, answers, true)
}

func outputPdf(filePath string, title string, items [][]string, answers string, printAnswer bool) {
	begin := time.Now()
	// 蓝色
	blueColor := getBlueColor()
	// 红色
	redColor := getRedColor()
	// 深灰色
	darkGrayColor := getDarkGrayColor()
	// 灰色
	grayColor := getGrayColor()

	// 实例化Pdf对象，指定纸张方向和大小
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	// 设置页面边距(左、上、右)单位是GE
	m.SetPageMargins(10, 5, 10)
	// 每行 12 GE

	// 加载字体
	m.AddUTF8Font("HarmonyOS", consts.Normal, "font/HarmonyOS_Sans_SC_Regular.ttf")
	m.AddUTF8Font("HarmonyOS", consts.Bold, "font/HarmonyOS_Sans_SC_Bold.ttf")
	m.AddUTF8Font("HarmonyOS", consts.Italic, "font/HarmonyOS_Sans_SC_Thin.ttf")
	m.AddUTF8Font("HarmonyOS", consts.BoldItalic, "font/HarmonyOS_Sans_SC_Black.ttf")
	m.SetDefaultFontFamily("HarmonyOS")

	m.SetAliasNbPages("{nb}")
	m.SetFirstPageNb(1)

	// 设置页眉
	m.RegisterHeader(func() {
		// 设置一行，行高10GE
		m.Row(12, func() {
			// 设置一列，列宽3GE，内容是 Logo 图片
			m.Col(3, func() {
				_ = m.FileImage("logo128.png", props.Rect{
					Percent: 100,
				})
			})

			// 打印试卷题目
			m.Col(6, func() {
				m.Text(title, props.Text{
					Style: consts.Bold,
					Align: consts.Center,
					Size:  22,
				})
			})

			// 设置页眉右侧文字
			m.Col(3, func() {
				// 第一行文字
				m.Text("KAKA 自动出题系统", props.Text{
					Size:        14,
					Align:       consts.Right,
					Extrapolate: false,
					Color:       blueColor,
				})
				// 第二行文字
				m.Text("难易度：简单", props.Text{
					Top:   10,
					Style: consts.Italic,
					Size:  10,
					Align: consts.Right,
					Color: redColor,
				})
			})
		})
		// 页眉的分割线
		m.Line(10.0,
			props.Line{
				Color: darkGrayColor,
			})
	})

	// 设定页脚
	m.RegisterFooter(func() {
		m.Row(0, func() {
			// 打印自家网址
			m.Col(4, func() {
				m.Text("https://kaba-kama-kaka.net", props.Text{
					Top:   10,
					Style: consts.BoldItalic,
					Size:  10,
					Align: consts.Left,
					Color: grayColor,
				})
			})
			// 打印页号
			m.Col(4, func() {
				m.Text(strconv.Itoa(m.GetCurrentPage())+"/{nb}", props.Text{
					Top:   10,
					Align: consts.Center,
					Size:  10,
				})
			})
			// 可以打印自家二维码
			m.ColSpace(4)
		})
	})

	// 打印题目
	// m.TableList([]string{"", "", ""}, items, props.TableList{
	// 	ContentProp: props.TableListContent{
	// 		GridSizes: []uint{6, 6},
	// 	},
	// 	Align:                  consts.Left,
	// 	VerticalContentPadding: 40,
	// })

	justAdded := false
	// 打印题目
	for _, line := range items {
		m.Row(60, func() {
			for _, item := range line {
				m.Col(6, func() {
					m.Text(item, props.Text{
						Align: consts.Left,
						Size:  14,
					})
				})
			}
		})
		// if (i+1)%4 == 0 {
		// 	m.AddPage()
		// 	justAdded = true
		// } else {
		// 	justAdded = false
		// }
	}

	if printAnswer && !justAdded {
		// 新增一页用于打印答案
		m.AddPage()
	}
	if printAnswer {
		// 打印试卷答案
		m.Row(10, func() {
			m.Col(12, func() {
				m.Text("答案：", props.Text{
					Style:       consts.Bold,
					Align:       consts.Left,
					Size:        14,
					Extrapolate: false, // true: 文字超出边界时，不折行。false：折行
				})
				m.Text(answers, props.Text{
					Top:         14,
					Style:       consts.Normal,
					Align:       consts.Left,
					Size:        14,
					Extrapolate: false, // true: 文字超出边界时，不折行。false：折行
				})
			})
		})
	}

	err := m.OutputFileAndClose(filePath)
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func TestPdf() {
	begin := time.Now()

	// 深灰色
	darkGrayColor := getDarkGrayColor()
	// 灰色
	grayColor := getGrayColor()
	// 白色
	whiteColor := color.NewWhite()
	// 蓝色
	blueColor := getBlueColor()
	// 红色
	redColor := getRedColor()
	// 表格标题
	header := getHeader()
	// 表格内容
	contents := getContents()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(3, func() {
				_ = m.FileImage("logo128.png", props.Rect{
					Center:  true,
					Percent: 100,
				})
			})

			m.ColSpace(6)

			m.Col(3, func() {
				m.Text("AnyCompany Name Inc. 851 Any Street Name, Suite 120, Any City, CA 45123.", props.Text{
					Size:        8,
					Align:       consts.Right,
					Extrapolate: false,
					Color:       redColor,
				})
				m.Text("Tel: 55 024 12345-1234", props.Text{
					Top:   12,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Right,
					Color: blueColor,
				})
				m.Text("www.mycompany.com", props.Text{
					Top:   15,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Right,
					Color: blueColor,
				})
			})
		})
	})
	////////////////////////////////////////
	m.RegisterFooter(func() {
		m.Row(20, func() {
			m.Col(12, func() {
				m.Text("Tel: 55 024 12345-1234", props.Text{
					Top:   13,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: blueColor,
				})
				m.Text("www.mycompany.com", props.Text{
					Top:   16,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: blueColor,
				})
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Invoice ABC123456789", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.SetBackgroundColor(darkGrayColor)

	m.Row(7, func() {
		m.Col(3, func() {
			m.Text("Transactions", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Center,
				Color: color.NewWhite(),
			})
		})
		m.ColSpace(9)
	})

	m.SetBackgroundColor(whiteColor)

	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 4, 2, 3},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 4, 2, 3},
		},
		Align:                consts.Center,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	m.Row(20, func() {
		m.ColSpace(7)
		m.Col(2, func() {
			m.Text("Total:", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			m.Text("R$ 2.567,00", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

	m.Row(15, func() {
		m.Col(6, func() {
			_ = m.Barcode("5123.151231.512314.1251251.123215", props.Barcode{
				Percent: 0,
				Proportion: props.Proportion{
					Width:  20,
					Height: 2,
				},
			})
			m.Text("5123.151231.512314.1251251.123215", props.Text{
				Top:    12,
				Family: "",
				Style:  consts.Bold,
				Size:   9,
				Align:  consts.Center,
			})
		})
		m.ColSpace(6)
	})

	err := m.OutputFileAndClose("pdfs/test.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

// 表格：标题
func getHeader() []string {
	return []string{"", "Product", "Quantity", "Price"}
}

// 表格：内容
func getContents() [][]string {
	return [][]string{
		{"", "Swamp", "12", "R$ 4,00"},
		{"", "Sorin, A Planeswalker", "4", "R$ 90,00"},
		{"", "Tassa", "4", "R$ 30,00"},
		{"", "Skinrender", "4", "R$ 9,00"},
		{"", "Island", "12", "R$ 4,00"},
		{"", "Mountain", "12", "R$ 4,00"},
		{"", "Plain", "12", "R$ 4,00"},
		{"", "Black Lotus", "1", "R$ 1.000,00"},
		{"", "Time Walk", "1", "R$ 1.000,00"},
		{"", "Emberclave", "4", "R$ 44,00"},
		{"", "Anax", "4", "R$ 32,00"},
		{"", "Murderous Rider", "4", "R$ 22,00"},
		{"", "Gray Merchant of Asphodel", "4", "R$ 2,00"},
		{"", "Ajani's Pridemate", "4", "R$ 2,00"},
		{"", "Renan, Chatuba", "4", "R$ 19,00"},
		{"", "Tymarett", "4", "R$ 13,00"},
		{"", "Doom Blade", "4", "R$ 5,00"},
		{"", "Dark Lord", "3", "R$ 7,00"},
		{"", "Memory of Thanatos", "3", "R$ 32,00"},
		{"", "Poring", "4", "R$ 1,00"},
		{"", "Deviling", "4", "R$ 99,00"},
		{"", "Seiya", "4", "R$ 45,00"},
		{"", "Harry Potter", "4", "R$ 62,00"},
		{"", "Goku", "4", "R$ 77,00"},
		{"", "Phreoni", "4", "R$ 22,00"},
		{"", "Katheryn High Wizard", "4", "R$ 25,00"},
		{"", "Lord Seyren", "4", "R$ 55,00"},
	}
}

// 颜色：深灰色
func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

// 颜色：灰色
func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

// 颜色：绿色
func getBlueColor() color.Color {
	return color.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

// 颜色：红色
func getRedColor() color.Color {
	return color.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}
