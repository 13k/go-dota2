package state

import (
	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// Dota2State is a snapshot of the client state at a point in time.
type Dota2State struct {
	// ConnectionStatus is the status of the connection to the GC.
	ConnectionStatus pb.GCConnectionStatus
	// Lobby is the current lobby object.
	Lobby *pb.CSODOTALobby
	// Party is the current party object.
	Party *pb.CSODOTAParty
	// PartyInvite is the active incoming party invite.
	PartyInvite *pb.CSODOTAPartyInvite
	// LastConnectionStatusUpdate is the last connection state update we received.
	LastConnectionStatusUpdate *pb.CMsgConnectionStatus
}

// ClearState clears everything.
func (s *Dota2State) ClearState() {
	*s = Dota2State{ConnectionStatus: pb.GCConnectionStatus_GCConnectionStatus_NO_SESSION}
}

// IsReady checks if the client is ready to receive requests.
func (s *Dota2State) IsReady() bool {
	return s.ConnectionStatus == pb.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION
}
