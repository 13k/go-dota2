package dota2

import (
	"google.golang.org/protobuf/proto"

	pb "github.com/13k/go-steam-resources/protobuf/dota2"
	"github.com/13k/go-steam/steamid"
)

// InviteLobbyMember attempts to invite a player to the current lobby.
func (d *Dota2) InviteLobbyMember(playerID steamid.SteamID) error {
	return d.write(uint32(pb.EGCBaseMsg_k_EMsgGCInviteToLobby), &pb.CMsgInviteToLobby{
		SteamId: proto.Uint64(playerID.Uint64()),
	})
}
