package xtime

import (
	"github.com/gogf/gf/v2/os/gtime"
	"strconv"
	"time"
)

// 获取某天的开始结束时间
func BeginOfDate(date time.Time) time.Time {
	timeStr := date.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t
}
func ToGTime(d time.Time) *gtime.Time {
	return gtime.New(d)
}

// 获取某天的开始结束时间
func BeginOfDateStr(date string) time.Time {
	parse, _ := time.Parse("2006-01-02", date)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", parse.Format("2006-01-02")+" 00:00:00", time.Local)
	return t
}

// 获取某天的最后结束时间
func EndOfDate(date time.Time) time.Time {
	timeStr := date.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	return t
}

// 获取某天的开始结束时间
func EndOfDateStr(date string) time.Time {
	parse, _ := time.Parse("2006-01-02", date)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", parse.Format("2006-01-02")+" 23:59:59", time.Local)
	return t
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func BeginDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return ZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func EndDateOfMonth(d time.Time) time.Time {
	return BeginDateOfMonth(d).AddDate(0, 1, -1)
}

//获取某一天的0点时间
func ZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 毫秒转时间
func MsToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Unix(0, msInt*int64(time.Millisecond))

	return tm, nil
}

func Century(year int) int {
	y := year/100 + 1
	return y
}
