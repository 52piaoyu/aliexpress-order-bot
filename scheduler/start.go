package scheduler

func StartScheduler(messageChan chan<- Message) {
	runningRoutine(checkUpdates, messageChan)
}
