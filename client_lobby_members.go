package dota2

import (
	pb "github.com/13k/go-steam-resources/protobuf/dota2"
	"github.com/13k/go-steam/steamid"
)

// InviteLobbyMember attempts to invite a player to the current lobby.
func (d *Dota2) InviteLobbyMember(playerID steamid.SteamID) {
	steamID := playerID.Uint64()
	d.write(uint32(pb.EGCBaseMsg_k_EMsgGCInviteToLobby), &pb.CMsgInviteToLobby{
		SteamId: &steamID,
	})
}
