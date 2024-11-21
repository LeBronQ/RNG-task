package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const (
	TypeNodeDelivery = "nodes:deliver"
)

type Node struct {
	ID int
	X  int
	Y  int
}

type NodeDeliveryPayload struct {
	Nodes []Node
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewNodesDeliveryTask(t []Node) (*asynq.Task, error) {
	payload, err := json.Marshal(NodeDeliveryPayload{Nodes: t})
	fmt.Print("p:", string(payload))
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeNodeDelivery, payload), nil
}
