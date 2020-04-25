package dota2

import (
	"context"

	"github.com/13k/go-dota2/cso"
	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// CreateLobby attempts to create a lobby with details.
func (d *Dota2) CreateLobby(details *pb.CMsgPracticeLobbySetDetails) error {
	return d.write(uint32(pb.EDOTAGCMsg_k_EMsgGCPracticeLobbyCreate), &pb.CMsgPracticeLobbyCreate{
		PassKey:      details.PassKey,
		LobbyDetails: details,
	})
}

// LeaveCreateLobby attempts to leave any current lobby and creates a new one.
func (d *Dota2) LeaveCreateLobby(
	ctx context.Context,
	details *pb.CMsgPracticeLobbySetDetails,
	destroyOldLobby bool,
) error {
	cacheCtr, err := d.cache.GetContainerForTypeID(uint32(cso.Lobby))
	if err != nil {
		return err
	}

	eventCh, eventCancel, err := cacheCtr.Subscribe()
	if err != nil {
		return err
	}
	defer eventCancel()

	var wasInNoLobby bool
	for {
		lobbyObj := cacheCtr.GetOne()
		if lobbyObj != nil {
			lob := lobbyObj.(*pb.CSODOTALobby)
			le := d.le.WithField("lobby-id", lob.GetLobbyId())
			if wasInNoLobby {
				le.Debug("successfully created lobby")
				return nil
			}

			le.Debug("attempting to leave lobby")
			if destroyOldLobby && lob.GetLeaderId() == d.client.SteamID().Uint64() {
				resp, err := d.DestroyLobby(ctx)
				if err != nil {
					return err
				}
				le.WithField("result", resp.GetResult().String()).Debug("destroy lobby result")
			}
			if lob.GetState() != pb.CSODOTALobby_UI {
				d.AbandonLobby()
			}
			d.LeaveLobby()
		} else {
			wasInNoLobby = true
			d.le.Debug("creating lobby")

			if err := d.CreateLobby(details); err != nil {
				return err
			}
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-eventCh:
		}
	}
}
