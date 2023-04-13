package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/oklog/ulid/v2"
)

type MessageType string

func (m MessageType) String() string {
	return string(m)
}

const (
	Generate   MessageType = "generate"
	GenerateOK MessageType = "generate_ok"
	Echo       MessageType = "echo"
	EchoOK     MessageType = "echo_ok"
)

type Body struct {
	Type      MessageType `json:"type,omitempty"`
	MessageID int         `json:"msg_id,omitempty"`
	ID        string      `json:"id,omitempty"`
	Echo      string      `json:"echo,omitempty"`
}

// TODO: this may actually not be needed as we have maelstrom.Message
type Message struct {
	Source      string `json:"src,omitempty"`
	Destination string `json:"dst,omitempty"`
	Body        Body   `json:"body,omitempty"`
	InReplyTo   int    `json:"in_reply_to,omitempty"`
}

func main() {
	node := maelstrom.NewNode()

	//for ulid generation, let's setup some needed entropy
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())

	node.Handle(Echo.String(), func(msg maelstrom.Message) error {
		// Unmarshal message body as loosely-typed map
		var body Body
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update message type
		body.Type = EchoOK

		// Echo back
		return node.Reply(msg, body)
	})

	node.Handle(Generate.String(), func(msg maelstrom.Message) error {
		var body Body
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body.Type = GenerateOK

		id, err := ulid.New(ms, entropy)
		if err != nil {
			return err
		}
		body.ID = id.String()

		return node.Reply(msg, body)
	})

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}
