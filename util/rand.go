package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// for i := 0; i < 10; i++ {
// 	value := rand.Intn(10)//Intn(10) 左閉右開區間 [0,10)
// 	fmt.Println(value)
// }

// 生成随机整数
func RandInt(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}

// 生成排除0的随机整数
func RandIntExcludeZero(min, max int) int {
	return RandIntExclude(min, max, 0)
}

// 生成排除某数的随机整数
func RandIntExclude(min, max int, exclude ...int) int {
	if Contains(exclude, min) && min == max {
		return exclude[0] + 1
	}

	for {
		value := RandInt(min, max)
		if !Contains(exclude, value) {
			return value
		}
	}
}

// 判断某个数是否存在数组中
func Contains(values []int, value int) bool {
	for _, i := range values {
		if i == value {
			return true
		}
	}
	return false
}

// 生成随机布尔值
func RandBool() bool {
	return rand.Intn(2) == 0
}
