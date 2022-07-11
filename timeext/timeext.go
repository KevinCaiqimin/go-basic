package timeext

import "caiqimin.tech/basic/mathext"
import "time"


func OffsetMs(time1, time2 time.Time, decimal int32) float64 {
	dura := time2.Sub(time1)
	sec := dura.Seconds()
	return mathext.KeepDecimal(sec * 1000, decimal)
}

func RandSleep(min, max int) {
	sleepTime := mathext.RandByMinMax(min, max)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
}