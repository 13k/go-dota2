package main

import (
	"strings"

	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// GetMessageEventName returns the event name for the message.
func GetMessageEventName(msg pb.EDOTAGCMsg) string {
	if over, ok := msgEventNameOverrides[msg]; ok {
		return over
	}

	msgName := msg.String()
	msgName = strings.TrimPrefix(msgName, "k_EMsg")
	msgName = strings.TrimPrefix(msgName, "GC")
	msgName = strings.TrimPrefix(msgName, "ToClient")
	msgName = strings.Replace(msgName, "_", "", -1)

	return msgName
}
