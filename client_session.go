package dota2

import (
	"context"

	"github.com/13k/go-dota2/events"
	"github.com/13k/go-dota2/state"
	pb "github.com/13k/go-steam-resources/protobuf/dota2"
	"github.com/13k/go-steam/protocol/gc"
)

// SetPlaying informs Steam we are playing / not playing Dota 2.
func (d *Dota2) SetPlaying(playing bool) {
	if playing {
		d.client.GC.SetGamesPlayed(AppID)
	} else {
		d.client.GC.SetGamesPlayed()
		_ = d.accessState(func(ns *state.Dota2State) (changed bool, err error) {
			ns.ClearState()
			return true, nil
		})
	}
}

// SayHello says hello to the Dota2 server, in an attempt to get a session.
func (d *Dota2) SayHello(haveCacheVersions ...*pb.CMsgSOCacheHaveVersion) {
	d.le.Debug("sending hello to GC")
	partnerAccType := pb.PartnerAccountType_PARTNER_NONE
	engine := pb.ESourceEngine_k_ESE_Source2
	var clientSessionNeed uint32 = 104
	d.write(uint32(pb.EGCBaseClientMsg_k_EMsgGCClientHello), &pb.CMsgClientHello{
		ClientLauncher:      &partnerAccType,
		Engine:              &engine,
		ClientSessionNeed:   &clientSessionNeed,
		SocacheHaveVersions: haveCacheVersions,
	})
}

// validateConnectionContext checks if the client is ready or not.
func (d *Dota2) validateConnectionContext() (context.Context, error) {
	d.connectionCtxMtx.Lock()
	defer d.connectionCtxMtx.Unlock()

	cctx := d.connectionCtx
	if cctx == nil {
		return nil, ErrNotReady
	}

	select {
	case <-cctx.Done():
		return nil, ErrNotReady
	default:
		return cctx, nil
	}
}

// handleClientWelcome handles an incoming client welcome event.
func (d *Dota2) handleClientWelcome(packet *gc.Packet) error {
	welcome := &pb.CMsgClientWelcome{}
	if err := d.unmarshalBody(packet, welcome); err != nil {
		return err
	}

	d.le.Debug("received GC welcome")
	for _, cache := range welcome.GetUptodateSubscribedCaches() {
		d.RequestCacheSubscriptionRefresh(cache.GetOwnerSoid())
	}

	for _, cache := range welcome.GetOutofdateSubscribedCaches() {
		if err := d.cache.HandleSubscribed(cache); err != nil {
			d.le.WithError(err).Warn("unable to handle welcome cache")
		}
	}

	d.setConnectionStatus(pb.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION, nil)
	d.emit(&events.ClientWelcomed{Welcome: welcome})
	return nil
}

// handleConnectionStatus handles the connection status update event.
func (d *Dota2) handleConnectionStatus(packet *gc.Packet) error {
	stat := &pb.CMsgConnectionStatus{}
	if err := d.unmarshalBody(packet, stat); err != nil {
		return err
	}

	if stat.Status == nil {
		return nil
	}

	d.setConnectionStatus(*stat.Status, stat)
	return nil
}

// setConnectionStatus sets the connection status, and emits an event.
// NOTE: do not call from inside accessState.
func (d *Dota2) setConnectionStatus(
	connStatus pb.GCConnectionStatus,
	update *pb.CMsgConnectionStatus,
) {
	_ = d.accessState(func(ns *state.Dota2State) (changed bool, err error) {
		if ns.ConnectionStatus == connStatus {
			return false, nil
		}

		oldState := ns.ConnectionStatus
		d.le.WithField("old", oldState.String()).
			WithField("new", connStatus.String()).
			Debug("connection status changed")
		d.emit(&events.GCConnectionStatusChanged{
			OldState: oldState,
			NewState: connStatus,
			Update:   update,
		})

		ns.ClearState() // every time the state changes, we lose the lobbies / etc
		ns.ConnectionStatus = connStatus
		d.connectionCtxMtx.Lock()
		if d.connectionCtxCancel != nil {
			d.connectionCtxCancel()
			d.connectionCtxCancel = nil
			d.connectionCtx = nil
		}
		if connStatus == pb.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION {
			d.connectionCtx, d.connectionCtxCancel = context.WithCancel(context.Background())
		}
		d.connectionCtxMtx.Unlock()
		return true, nil
	})
}
