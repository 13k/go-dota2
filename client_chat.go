package dota2

import (
	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// SendChannelMessage attempts to send a message in a channel, text-only.
// Use SendChatMessage for more fine-grained control.
func (d *Dota2) SendChannelMessage(channelID uint64, message string) error {
	return d.write(uint32(pb.EDOTAGCMsg_k_EMsgGCChatMessage), &pb.CMsgDOTAChatMessage{
		ChannelId: &channelID,
		Text:      &message,
	})
}
