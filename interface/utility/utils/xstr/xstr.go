package xstr

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CheckEasyUname(uname string) error {
	num := len(uname)
	if num < 4 || num > 16 {
		return errors.New("用户名为4到16位字符")
	}
	return nil
}

// 截取小数点后面的数字
// index 1 代表 小数点后第一位， 2代表小数点后第二位
func SubPricePointNum(num float64, index int) int64 {
	numStr := fmt.Sprintf("%f", num)
	i := strings.Index(numStr, ".")
	flagStr := numStr[i+index : i+index+1]
	parseInt, err := strconv.ParseInt(flagStr, 10, 32)
	if err != nil {
		panic(err)
	}
	return parseInt
}
func Like(str string) string {
	return fmt.Sprint("%", str, "%")
}
func LikePre(str string) string {
	return fmt.Sprint(str, "%")
}
func LikeSuffix(str string) string {
	return fmt.Sprint("%", str)
}
func ContainsAny(s string, substr ...string) bool {
	for _, item := range substr {
		if strings.Contains(s, item) {
			return true
		}
	}
	return false
}

func RandomUnameWithDate() int64 {
	return time.Now().Unix()
}
