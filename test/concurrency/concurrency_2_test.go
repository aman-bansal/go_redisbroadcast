package concurrency

import (
"fmt"
"github.com/aman-bansal/go_redisbroadcast"
"testing"
)

func TestConcurrent_2(t *testing.T) {
	err := go_redisbroadcast.Init()
	if err != nil {
		t.Error(err)
		return
	}

	err = go_redisbroadcast.Register("EVENT_TYPE_1", new(MessageProcess2))
	if err != nil {
		t.Error(err)
		return
	}

	err = go_redisbroadcast.Publish("EVENT_TYPE_1", go_redisbroadcast.Message{
		MessageId:   "1234567890",
		MessageText: "hello aman",
	})
	if err != nil {
		t.Error(err)
		return
	}

	go_redisbroadcast.Close()
}

type MessageProcess2 struct {}

func (m *MessageProcess2) Process(eventType string, message go_redisbroadcast.Message) {
	fmt.Printf("In Process Message 2 %s:%+v", eventType, message)
}
