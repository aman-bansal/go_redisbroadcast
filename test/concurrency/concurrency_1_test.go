package concurrency

import (
	"fmt"
	"github.com/aman-bansal/go_redisbroadcast"
	"testing"
	"time"
)

func TestConcurrent_1(t *testing.T) {
	err := go_redisbroadcast.Init()
	if err != nil {
		t.Error(err)
		return
	}

	err = go_redisbroadcast.Register("EVENT_TYPE_1", new(MessageProcess))
	if err != nil {
		t.Error(err)
		return
	}


	time.Sleep(time.Second * 100)
	go_redisbroadcast.Close()

	//err = go_redisbroadcast.Publish("EVENT_TYPE_1", go_redisbroadcast.Message{
	//	MessageId:   "1234567890",
	//	MessageText: "hello aman",
	//})
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
}

type MessageProcess struct {}

func (m *MessageProcess) Process(eventType string, message go_redisbroadcast.Message) {
	fmt.Printf("In Process Message 1 %s:%+v", eventType, message)
}
