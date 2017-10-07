package serve_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/yizenghui/sda/code"
)

func Test_String2Int(t *testing.T) {
	i, err := strconv.Atoi("3200000")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}

func Test_String2Int64(t *testing.T) {

	i64, err := strconv.ParseInt("3200000", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i64)
}

func Test_TimeString2Int64(t *testing.T) {

	i64, err := strconv.ParseInt("1507345151", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(i64)

}

func Test_String2Int32(t *testing.T) {

	i64, err := strconv.ParseInt("s3200000", 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	i32 := int32(i64)
	fmt.Println(i32)
}

func Test_StringMateInt(t *testing.T) {

	str := "ss_123asd"
	i64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		// 这里面用正则匹配出整数
		istr := code.FindString(`(?P<int>\d+)`, str, "int")
		i64, _ = strconv.ParseInt(istr, 10, 64)
	}
	fmt.Println(i64)

}
