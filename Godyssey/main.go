package main

import (
	"encoding/json"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
)

func main() {
	node := maelstrom.NewNode()

	// TODO: lets standardize the following below and create structs, types, and enums

	node.Handle("echo", func(msg maelstrom.Message) error {
		// Unmarshal message body as loosely-typed map
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update message type
		body["type"] = "echo_ok"

		// Echo back
		return node.Reply(msg, body)
	})

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}
