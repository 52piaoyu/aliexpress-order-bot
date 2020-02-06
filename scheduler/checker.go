package scheduler

func checkUpdates(messageChan chan<- Message) {
	var message = Message{0, "You nigga"}
	messageChan <- message
}
