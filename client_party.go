package dota2

import (
	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// LeaveParty attempts to leave the current party.
func (d *Dota2) LeaveParty() {
	d.write(uint32(pb.EGCBaseMsg_k_EMsgGCLeaveParty), &pb.CMsgLeaveParty{})
}
