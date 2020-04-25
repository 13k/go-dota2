package events

import (
	"google.golang.org/protobuf/proto"

	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// ChatMessage is emitted when a chat message is observed.
type ChatMessage struct {
	pb.CMsgDOTAChatMessage
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ChatMessage) GetDotaEventMsgID() pb.EDOTAGCMsg {
	return pb.EDOTAGCMsg_k_EMsgGCChatMessage
}

// GetEventBody returns the event body.
func (e *ChatMessage) GetEventBody() proto.Message {
	return &e.CMsgDOTAChatMessage
}

// GetEventName returns the chat message event name.
func (e *ChatMessage) GetEventName() string {
	return "ChatMessage"
}

// JoinedChatChannel is emitted when we join a chat channel.
type JoinedChatChannel struct {
	pb.CMsgDOTAJoinChatChannelResponse
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *JoinedChatChannel) GetDotaEventMsgID() pb.EDOTAGCMsg {
	return pb.EDOTAGCMsg_k_EMsgGCJoinChatChannelResponse
}

// GetEventBody returns the event body.
func (e *JoinedChatChannel) GetEventBody() proto.Message {
	return &e.CMsgDOTAJoinChatChannelResponse
}

// GetEventName returns the chat message event name.
func (e *JoinedChatChannel) GetEventName() string {
	return "JoinedChatChannel"
}
