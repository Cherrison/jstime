package jstime

import "time"

// 指定周期执行一个函数
func SetInterval(d time.Duration, f func()) * time.Ticker {
	ticker := time.NewTicker(d)
	go func() {
		for range ticker.C {
			f()
		}
	}()
	return ticker
}

func SetTimeout(d time.Duration, f func()) * time.Timer {
	return time.AfterFunc(d , f)
}
