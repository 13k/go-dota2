package dota2

import (
	pb "github.com/13k/go-steam-resources/protobuf/dota2"
	"github.com/13k/go-steam/protocol/gc"
)

// RequestCacheSubscriptionRefresh requests a subscription refresh for a specific cache ID.
func (d *Dota2) RequestCacheSubscriptionRefresh(ownerSoid *pb.CMsgSOIDOwner) {
	d.write(uint32(pb.ESOMsg_k_ESOMsg_CacheSubscriptionRefresh), &pb.CMsgSOCacheSubscriptionRefresh{
		OwnerSoid: ownerSoid,
	})
}

// handleCacheSubscribed handles a CacheSubscribed packet.
func (d *Dota2) handleCacheSubscribed(packet *gc.Packet) error {
	sub := &pb.CMsgSOCacheSubscribed{}
	if err := d.unmarshalBody(packet, sub); err != nil {
		return err
	}

	if err := d.cache.HandleSubscribed(sub); err != nil {
		d.le.WithError(err).Debug("unhandled cache issue (ignore)")
	}

	return nil
}

// handleCacheUnsubscribed handles a CacheUnsubscribed packet.
func (d *Dota2) handleCacheUnsubscribed(packet *gc.Packet) error {
	sub := &pb.CMsgSOCacheUnsubscribed{}
	if err := d.unmarshalBody(packet, sub); err != nil {
		return err
	}

	if err := d.cache.HandleUnsubscribed(sub); err != nil {
		d.le.WithError(err).Debug("unhandled cache issue (ignore)")
	}

	return nil
}

// handleCacheUpdateMultiple handles when one or more object(s) in a cache is/are updated.
func (d *Dota2) handleCacheUpdateMultiple(packet *gc.Packet) error {
	sub := &pb.CMsgSOMultipleObjects{}
	if err := d.unmarshalBody(packet, sub); err != nil {
		return err
	}

	if err := d.cache.HandleUpdateMultiple(sub); err != nil {
		d.le.WithError(err).Debug("unhandled cache issue (ignore)")
	}

	return nil
}

// handleCacheDestroy handles when an object in a cache is destroyed.
func (d *Dota2) handleCacheDestroy(packet *gc.Packet) error {
	sub := &pb.CMsgSOSingleObject{}
	if err := d.unmarshalBody(packet, sub); err != nil {
		return err
	}

	if err := d.cache.HandleDestroy(sub); err != nil {
		d.le.WithError(err).Debug("unhandled cache issue (ignore)")
	}

	return nil
}
