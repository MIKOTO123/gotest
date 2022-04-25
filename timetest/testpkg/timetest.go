package testpkg

import (
	"fmt"
	"time"
)

func TimeTest() {
	now := time.Now()
	fmt.Printf("t: %T\n", now)
	fmt.Printf("t: %v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T,%T,%T,%T,%T,%T\n", year, month, day, hour, minute, second)
}

func TimeTest2() {
	now := time.Now()
	// 当前时间戳 TimeStamp type:int64, TimeStamp:1606832965
	fmt.Printf("TimeStamp type:%T, TimeStamp:%v", now.Unix(), now.Unix())
}

func TimeTest3() {
	t := time.Now()
	fmt.Println(t.Format("2006/01/02 15:04")) //这个2006/01/02 15:04的格式是固定的值
}

func TimeTest4(h, m, s, mls, msc, ns time.Duration) {
	now := time.Now()
	fmt.Println(now.Add(time.Hour*h + time.Minute*m + time.Second*s + time.Millisecond*mls + time.Microsecond*msc + time.Nanosecond*ns))
}

func TimeTest5() {
	now := time.Now()
	fmt.Println(now.AddDate(1, 5, 3))
}
func TimeTest6() {

	fmt.Printf("%v\n", time.Now().Format(time.RFC3339))
	parse, _ := time.Parse("2006-01-02 15:04:05", "2022-04-25 20:49:36") //默认是UTC时区
	fmt.Printf("%T\n", parse)
	fmt.Printf("%v\n", parse)
	location, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-04-25 20:49:36", time.Local)
	fmt.Printf("%T\n", location)
	fmt.Printf("%v\n", location)

}
