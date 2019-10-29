package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("x = %.6f\n", 21312421.213123)
	fmt.Printf("x = %.f\n", 21312421.213123)

	// FormatFloat 将浮点数 f 转换为字符串值
	// f：要转换的浮点数
	// fmt：格式标记（b、e、E、f、g、G）
	// prec：精度（数字部分的长度，不包括指数部分）
	// bitSize：指定浮点类型（32:float32、64:float64）
	//
	// 格式标记：
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
	//
	// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
	// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
	f := 21312421.213123
	s := strconv.FormatFloat(f, 'f', 0, 64)
	fmt.Println(s)
}
