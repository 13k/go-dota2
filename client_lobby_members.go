package dota2

import (
	"github.com/13k/go-steam/steamid"

	bgcm "github.com/13k/go-dota2/protocol"
)

// InviteLobbyMember attempts to invite a player to the current lobby.
func (d *Dota2) InviteLobbyMember(playerID steamid.SteamID) {
	steamID := playerID.Uint64()
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCInviteToLobby), &bgcm.CMsgInviteToLobby{
		SteamId: &steamID,
	})
}
