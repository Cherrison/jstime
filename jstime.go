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
/*
SetTimeout waits for the duration to elapse and then calls f in its own goroutine. It returns a Timer that can be used to cancel the call using its Stop method.
 */
func SetTimeout(d time.Duration, f func()) * time.Timer {
	return time.AfterFunc(d , f)
}
