package timerModel

import (
	"container/list"
)

type TimerManagerInstance struct {
	timerList *list.List  // pointer to timer
	counter int64
	owner string
}

func GetTimerManager() (*TimerManagerInstance) {
	instance := new(TimerManagerInstance);
	instance.timerList = new(list.List);
	
	return instance;
}

func (self *TimerManagerInstance) SetOwnerName (name string) {
	self.owner = name;
}

// 在外面產生一個timer，
// 調用此function，
// 該timer便會進入timerList，被manager服務
func (self *TimerManagerInstance) RegisteTimer (t *Timer) {
	self.timerList.PushBack(t);
	self.counter++;
}

// 幫忙產生一個 timer並註冊到 list 裡面
func (self *TimerManagerInstance) CreateAndRegisteTimer (interval int64) (t *Timer){
	// 產生 timer
	timer := new(Timer);
	timer.Interval = interval;
	timer.OnCounting = false;
	// 註冊 timer
	self.RegisteTimer(timer);
	return timer;
}

