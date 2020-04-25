package events

import (
	pb "github.com/13k/go-steam-resources/protobuf/dota2"
	"github.com/13k/go-steam/protocol/gc"
)

// GCConnectionStatusChanged is emitted when the client connection state is updated.
type GCConnectionStatusChanged struct {
	// OldState contains the old connection status.
	OldState pb.GCConnectionStatus
	// NewState contains the new connection status.
	NewState pb.GCConnectionStatus
	// Update contains the message from the server that triggered this change, may be nil.
	Update *pb.CMsgConnectionStatus
}

// ClientWelcomed is emitted when the client receives the GC welcome
type ClientWelcomed struct {
	// Welcome is the welcome message from the GC.
	Welcome *pb.CMsgClientWelcome
}

// UnhandledGCPacket is called when the client ignores an unhandled packet.
type UnhandledGCPacket struct {
	// Packet is the unhandled packet.
	Packet *gc.Packet
}
