package go_redisbroadcast

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

type Message struct {
	MessageId   string `json:"message_id"`
	MessageText string `json:"message_text"`
}

var redisClient *redis.Client
var eventVsPubSub = make(map[string]*redis.PubSub)
var err error

var mutex sync.Mutex

func Close() {
	_ = redisClient.Close()
}
func Init() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		//Password: "", // no password set
		//DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	return nil
}

type Process interface {
	Process(eventType string, message Message)
}

func Register(eventType string, process Process) error {
	if eventVsPubSub[eventType] != nil {
		return nil
	}
	pubSub := redisClient.Subscribe(eventType)
	_, err := pubSub.Receive()
	if err != nil {
		return err
	}

	mutex.Lock()
	defer  mutex.Unlock()
	eventVsPubSub[eventType] = pubSub
	go func() { _ = listen(pubSub.Channel(), process) }()
	return nil
}

func listen(channel <-chan *redis.Message, process Process) error {
	for {
		v := <-channel
		msg := new(Message)
		_ = json.Unmarshal([]byte(v.Payload), msg)
		process.Process(v.Channel, * msg)
		//fmt.Println(v.Channel, v.Payload)
	}
}

func Publish(eventType string, message Message) error {
	bytes, _ := json.Marshal(message)
	err = redisClient.Publish(eventType, string(bytes)).Err()
	if err != nil {
		return err
	}
	return nil
}
