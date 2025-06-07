package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	values := make([]float64, 0)

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		res := make(map[string]string)
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		v := body["message"].(float64)
		values = append(values, v)
		res["type"] = "broadcast_ok"

		return n.Reply(msg, res)
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		var body map[string]any
		res := make(map[string]any)
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		res["type"] = "read_ok"
		res["messages"] = values

		return n.Reply(msg, res)
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		var body map[string]any
		res := make(map[string]string)
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		res["type"] = "topology_ok"

		return n.Reply(msg, res)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
