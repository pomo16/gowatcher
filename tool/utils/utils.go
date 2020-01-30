package utils

import (
	"gowatcher/go_analyze/consts"
	"strings"
	"time"
)

//GetTimeStampByTimeStr 通过时间字符串获取时间戳
func GetTimeStampByTimeStr(base string) int64 {
	tt, _ := time.ParseInLocation(consts.TimeStr, base, time.Local)
	return tt.Unix()
}

//SplitTitleAndContent 分离标题和内容
func SplitTitleAndContent(base string) (string, string) {
	s := strings.Split(base, "|")
	if len(s) == 2 {
		return s[0], s[1]
	}

	return s[0], s[0]
}
