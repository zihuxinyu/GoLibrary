package Library

import (
	"time"
	"github.com/astaxie/beego"
	"fmt"
)
const TIME_LAYOUT_OFTEN = "2006-01-02 15:04:05"
// 解析常用的日期时间格式：2014-01-11 16:18:00，东八区
func TimeParseOften(value string) (time.Time, error) {
	local, _ := time.LoadLocation("Local")
	return time.ParseInLocation(TIME_LAYOUT_OFTEN, value, local)
}

//返回当前时区的当前时间
func TimeLocal() ( time.Time) {
	stime:="2006-01-02 15:04:05 -07:00 "
	datastring:=beego.DateFormat(time.Now(),stime)
	rtime,_:=beego.DateParse(datastring,stime)
	fmt.Println(datastring,rtime)
	return rtime
}
