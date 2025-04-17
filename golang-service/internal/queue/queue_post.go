package queue

import (
	"encoding/json"
	"os"

	"github.com/BRIZINGR007/go-service-utils/clients/sqs_client"
	"github.com/BRIZINGR007/go-service-utils/structs"
)

func CodeAssistQueuePostMessageNF(event string, payload any, context map[string]string) error {
	queueURL := os.Getenv("SQS_QUEUE_URL")
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	message_payload := &structs.MessageBody{
		Context: context,
		Event:   event,
		Payload: payloadBytes,
	}
	return sqs_client.PostMessageNonFIFO(queueURL, message_payload)

}
