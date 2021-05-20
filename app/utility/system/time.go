package system

import "time"

// GetCurrentDate 获取当前的时间 - 字符串
func GetCurrentDate() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// GetCurrentUnix 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// GetCurrentMilliUnix 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// GetCurrentNanoUnix 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}

// FormatDate 时间戳转时间格式
func FormatDate(timestamp int64) string {
	layout := "2006-01-02 15:04:05"
	return time.Unix(timestamp, 0).Format(layout)
}



// Duration 记录时间点
type Duration interface {
	Observe(handle func(time.Time, time.Time, time.Duration))
}

type duration struct {
	start time.Time
	end   time.Time
}

// Observe 查看时间详细信息
func (d *duration) Observe(handle func(time.Time, time.Time, time.Duration)) {
	delta := d.end.Sub(d.start)
	handle(d.start, d.end, delta)
}

// TimeLog 记录执行时间
func TimeLog(f func()) Duration {
	d := &duration{
		start: time.Now(),
	}
	f()
	d.end = time.Now()
	return d
}