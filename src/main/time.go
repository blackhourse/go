package main

import (
	"fmt"
	"time"
)

func main() {
	//getNow()
	//timeTrans()
	//times()
	//timeTick()
	//timeformatter()
	timeprase()
}

func getNow() {

	now := time.Now()
	fmt.Println("now time", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//	 时间戳
	unix := now.Unix()     // 时间戳
	nano := now.UnixNano() // 纳秒时间戳
	fmt.Printf("时间戳：%v\n", unix)
	fmt.Printf("纳秒时间戳：%v\n", nano)

}

// 将时间戳 转换为时间格式
func timeTrans() {
	now := time.Now()
	timestamp := now.Unix()
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func times() {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(-time.Hour)
	fmt.Println(later)
}

// 定时器
func timeTick() {
	fmt.Println(time.Second)
	tick := time.Tick(time.Second)
	for i := range tick {
		fmt.Println(i)
	}
}

// 时间格式化  时间类型有一个自带的 Format 方法进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S 而是使用Go语言的诞生时间 2006 年 1 月 2 号 15 点 04 分 05 秒。
func timeformatter() {
	now := time.Now()
	fmt.Println(now.Format("2006/1/2 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
// 解析字符串格式的格式
func timeprase()  {
	var layout string = "2006-01-02 15:04:05"
	var timeStr string = "2019-12-12 15:22:12"
	timeObj1, _ := time.Parse(layout, timeStr)
	fmt.Println(timeObj1)
	location, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println(location)


}