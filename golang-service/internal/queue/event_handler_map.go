package queue

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/consumers"
	"github.com/BRIZINGR007/go-service-utils/clients/sqs_client"
)

type Event1Payload struct {
	GithubURL string `json:"github_url"`
	AccessKey string `json:"access_key"`
}

type Event2Payload struct {
	ChatId string `json:"chat_id"`
	Chat   string `json:"chat"`
}

func handleEvent1(payload Event1Payload) error {
	fmt.Println("Waiting  for 10 seconds  ...")
	time.Sleep(1 * time.Second)
	fmt.Println("Handling event 1 with payload :", payload)
	return nil
}

func wrap[T any](handler func(T) error) func(json.RawMessage) error {
	return func(payload json.RawMessage) error {
		var data T

		// If T is []byte, handle it differently
		if _, ok := any(*(new(T))).([]byte); ok {
			// For []byte type, copy the raw JSON directly
			dataPtr := any(&data).(*([]byte))
			*dataPtr = payload
			return handler(data)
		}

		// Regular unmarshaling for other types
		if err := json.Unmarshal(payload, &data); err != nil {
			return err
		}
		return handler(data)
	}
}

var EventHandlers = map[string]func(json.RawMessage) error{
	"event1":          wrap(handleEvent1),
	"add-codecontext": wrap(consumers.AddCodeContext),
}

func InitEventProcessor() {
	processor := sqs_client.GetEventProcessor()
	processor.EventHandlers = EventHandlers
}
