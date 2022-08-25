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

	for i := 0; i < 10; i++ {
		e := util.NewEquationOneOperand(util.SideOne, util.NeedUncertain, 2, 2, 1, 1)

		fmt.Printf("%d. 答案：%d\t\t公式：%s\n", i+1, e.Value, e.String())
	}
}
