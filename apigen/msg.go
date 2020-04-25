package main

import (
	"sort"

	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// IsValidMsg checks if the message is valid.
func IsValidMsg(msg pb.EDOTAGCMsg) bool {
	_, ok := pb.EDOTAGCMsg_name[int32(msg)]
	return ok && msg > pb.EDOTAGCMsg_k_EMsgGCDOTABase
}

func getSortedMsgIDs() []pb.EDOTAGCMsg {
	var msgIds []pb.EDOTAGCMsg

	for msgIDNum := range pb.EDOTAGCMsg_name {
		msgID := pb.EDOTAGCMsg(msgIDNum)
		msgIds = append(msgIds, msgID)
	}

	sort.Slice(msgIds, func(i int, j int) bool {
		return msgIds[i] < msgIds[j]
	})

	return msgIds
}
