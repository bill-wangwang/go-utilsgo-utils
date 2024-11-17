package main

import (
	"fmt"
	"utils"
)

func main() {

	//日期转时间戳
	fmt.Println(utils.DateToTimestamp("2019-01-01"))
	fmt.Println(utils.DateTimeToTimestamp("2019-01-01 00:00:00"))

	//时间戳转日期
	fmt.Println(utils.TimestampToDate(1546300800))
	fmt.Println(utils.TimestampToDateTime(1546300800))

	//获取空闲端口
	port, _ := utils.GetFreePort()
	fmt.Println(port)
}
