package cso

import (
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// CSOType is a shared object type identifier.
//go:generate stringer -type=CSOType
type CSOType int32 //nolint: golint

const (
	// EconItem is an economy item.
	EconItem CSOType = 1
	// ItemRecipe is an item recipe.
	ItemRecipe = 5
	// EconGameAccountClient is a economy game account client..
	EconGameAccountClient = 7
	// SelectedItemPreset is a selected item preset.
	SelectedItemPreset = 35
	// ItemPresetInstance is a instance of an item preset.
	ItemPresetInstance = 36
	// DropRateBonus is an active drop rate bonus.
	DropRateBonus = 38
	// LeagueViewPass is a pass to view a league ticket.
	LeagueViewPass = 39
	// EventTicket is a ticket to an event.
	EventTicket = 40
	// ItemTournamentPassport is an item representing a tournament passport.
	ItemTournamentPassport = 42
	// GameAccountClient is the DOTA game account for a client.
	GameAccountClient = 2002
	// Party is a Dota 2 party.
	Party = 2003
	// Lobby is a Dota 2 lobby.
	Lobby = 2004
	// PartyInvite is an invite to a party.
	PartyInvite = 2006
	// GameHeroFavorites are game hero favorites.
	GameHeroFavorites = 2007
	// MapLocationState is the minimap location state.
	MapLocationState = 2008
	// Tournament represents a tournament.
	Tournament = 2009
	// PlayerChallenge represents a player challenge.
	PlayerChallenge = 2010
	// LobbyInvite is an invitation to a lobby.
	LobbyInvite = 2011
)

// csoTypeCtors links type IDs to constructors.
var csoTypeCtors = map[CSOType]func() proto.Message{
	EconItem: func() proto.Message {
		return &pb.CSOEconItem{}
	},
	GameAccountClient: func() proto.Message {
		return &pb.CSODOTAGameAccountClient{}
	},
	Party: func() proto.Message {
		return &pb.CSODOTAParty{}
	},
	Lobby: func() proto.Message {
		return &pb.CSODOTALobby{}
	},
	PartyInvite: func() proto.Message {
		return &pb.CSODOTAPartyInvite{}
	},
	GameHeroFavorites: func() proto.Message {
		return &pb.CSODOTAGameHeroFavorites{}
	},
	MapLocationState: func() proto.Message {
		return &pb.CSODOTAMapLocationState{}
	},
	PlayerChallenge: func() proto.Message {
		return &pb.CSODOTAPlayerChallenge{}
	},
	LobbyInvite: func() proto.Message {
		return &pb.CSODOTALobbyInvite{}
	},
	LeagueViewPass: func() proto.Message {
		return &pb.CSOEconItemLeagueViewPass{}
	},
	DropRateBonus: func() proto.Message {
		return &pb.CSOEconItemDropRateBonus{}
	},
}

// NewSharedObject builds a new shared object from a type ID.
func NewSharedObject(typ CSOType) (proto.Message, error) {
	ctor, ok := csoTypeCtors[typ]
	if !ok {
		return nil, errors.Errorf("unknown shared object type id: %d", typ)
	}

	return ctor(), nil
}
