package scheduler

import (
	"fmt"
	"time"
)

const intervalPeriod time.Duration = 24 * time.Hour

const hourToTick int = 21
const minuteToTick int = 55
const secondToTick int = 30

type jobTicker struct {
	timer *time.Timer
}

func runningRoutine(function func(chan<- Message), messageChan chan<- Message) {
	jobTicker := &jobTicker{}
	jobTicker.updateTimer()
	for {
		<-jobTicker.timer.C
		function(messageChan)
		fmt.Println(time.Now(), "- just ticked")
		jobTicker.updateTimer()
	}
}

func (t *jobTicker) updateTimer() {
	nextTick := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), hourToTick, minuteToTick, secondToTick, 0, time.Local)
	if !nextTick.After(time.Now()) {
		nextTick = nextTick.Add(intervalPeriod)
	}
	fmt.Println(nextTick, "- next tick")
	diff := nextTick.Sub(time.Now())
	if t.timer == nil {
		t.timer = time.NewTimer(diff)
	} else {
		t.timer.Reset(diff)
	}
}
