package events

import (
	"google.golang.org/protobuf/proto"

	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// Event is a DOTA event.
type Event interface {
	// GetDotaEventMsgID returns the DOTA event message ID.
	GetDotaEventMsgID() pb.EDOTAGCMsg
	// GetEventBody event body.
	GetEventBody() proto.Message
	// GetEventName returns the event name.
	GetEventName() string
}
