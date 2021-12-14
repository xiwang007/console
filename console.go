package console

import (
	"fmt"
	"time"
)

const Version = 1.0

// 打印日志
func Log(args ...interface{}) {
	var newArgs []interface{} = make([]interface{}, 0)
	//1. 获取当前时间
	now := time.Now()
	//格式化日期时间 2021/12/08 - 17:02:01
	dateStr := fmt.Sprintf("%04d/%02d/%02d - %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	newArgs = append(newArgs, dateStr)
	newArgs = append(newArgs, args...)
	fmt.Println(newArgs...)
}
