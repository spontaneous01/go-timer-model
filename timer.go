package timerModel

import (
	"time"
)

type Timer struct {
	Interval int64
	StartTime int64
	OnCounting bool
}

func (self *Timer) Start () {
	self.StartTime = getTimeMillionSecond();
	self.OnCounting = true;
}

func (self *Timer) Stop () {
	self.OnCounting = false;
}

func (self *Timer) SetInterval (interval int64) {
	self.Interval = interval;
}

func (self *Timer) SetIntervalAndStart (interval int64) {
	self.Interval = interval;
	self.Start();
}

func (self *Timer) IsCounting () (bool) {
	return self.OnCounting;
}

// 是否已經到達 timeout 標準
func (self *Timer) IsTimeout () (bool) {
	if cmf.GetTimeMillionSecond() - self.StartTime >= self.Interval {
		return true
	}
	return false
}

// 複製
func (self *Timer) Clone() (*Timer) {
	newTimer := new(Timer)
	newTimer.Interval = self.Interval
	newTimer.StartTime = self.StartTime
	newTimer.OnCounting = false
	return newTimer
}

//
func getTimeMillionSecond () int64 {
	return (time.Now().UnixNano() / 1e6);
}