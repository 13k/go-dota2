package dota2

import (
	"github.com/13k/go-dota2/events"
	"github.com/13k/go-steam/protocol/gc"
)

// handleMatchSignedOut handles an incoming steam datagram ticket.
func (d *Dota2) handleMatchSignedOut(packet *gc.Packet) error { //nolint: unused
	ev := &events.MatchSignedOut{}
	if err := d.unmarshalBody(packet, &ev.CMsgGCToClientMatchSignedOut); err != nil {
		return err
	}

	d.emit(ev)
	return nil
}
