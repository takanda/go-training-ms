package types

import (
	"fmt"
	"strconv"
)

func ShowBasicType() {
	// runeはシングルクォートで機能する。ダブルクオートの場合は文字列がそのまま出力される
	// シングルクォートは一つの文字にのみ使用し、それ以外はダブルクオートで囲む
	// rune := 'G'
	// fmt.Println(rune)
	
	// 指定したビット長を超えた場合はoverflowエラーが発生
	// var integer32 int32 = 2147483648
	// fmt.Println(integer32)
	
	// uintで負の値を代入した場合はoverflowエラーが発生
	// var integer uint = -10
	// fmt.Println(integer)

	// 規定値
	// var defaultInt int
	// var defaultFloat32 float32
	// var defaultFloat64 float64
	// var defaultBool bool
	// var defaultString string
	// fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

	// 型の変換
	// var integer16 int16 = 127
	// var integer32 int32 = 32767
	// fmt.Println(int32(integer16) + integer32)

	i, j := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, j, s)
}