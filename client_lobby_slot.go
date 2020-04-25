package dota2

import (
	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// JoinLobbyTeam switches team in a lobby.
func (d *Dota2) JoinLobbyTeam(team pb.DOTA_GC_TEAM, slot uint32) {
	d.write(uint32(pb.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot), &pb.CMsgPracticeLobbySetTeamSlot{
		Team: &team,
		Slot: &slot,
	})
}

// SetLobbySlotBotDifficulty sets the difficulty of a slot to a given bot difficulty.
func (d *Dota2) SetLobbySlotBotDifficulty(team pb.DOTA_GC_TEAM, slot uint32, botDifficulty pb.DOTABotDifficulty) {
	d.write(uint32(pb.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot), &pb.CMsgPracticeLobbySetTeamSlot{
		Team:          &team,
		Slot:          &slot,
		BotDifficulty: &botDifficulty,
	})
}
