package main

import (
	"google.golang.org/protobuf/proto"

	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

// msgSenderOverrides overrides the heuristic-generated sender parties for each message
// Most of the MsgSenderNone messages are not GC<->Client messages.
var msgSenderOverrides = map[pb.EDOTAGCMsg]MsgSender{
	pb.EDOTAGCMsg_k_EMsgGCLiveScoreboardUpdate:      MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCPlayerReports:             MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCPlayerFailedToConnect:     MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCGCToLANServerRelayConnect: MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCGCToRelayConnect:          MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCGCToRelayConnectresponse:  MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCSuggestTeamMatchmaking:    MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgClientsRejoinChatChannels: MsgSenderClient,

	pb.EDOTAGCMsg_k_EMsgGCConsumeFantasyTicket:        MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCConsumeFantasyTicketFailure: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCPlayerHeroesFavoritesAdd:    MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCPlayerHeroesFavoritesRemove: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCRewardDiretidePrizes: MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCRewardTutorialPrizes: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCGeneralResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel: MsgSenderGC,
	pb.EDOTAGCMsg_k_EMsgGCOtherLeftChannel:   MsgSenderGC,

	pb.EDOTAGCMsg_k_EMsgGCTeamData: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGC_TournamentItemEvent:         MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGC_TournamentItemEventResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgDOTAWeekendTourneySchedule: MsgSenderGC,

	pb.EDOTAGCMsg_k_EMsgGCMatchHistoryList: MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCGetRecentMatches: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCPracticeLobbyList: MsgSenderClient,

	pb.EDOTAGCMsg_k_EMsgGCInitialQuestionnaireResponse: MsgSenderClient,

	pb.EDOTAGCMsg_k_EMsgGCPracticeLobbyResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame: MsgSenderClient,
	pb.EDOTAGCMsg_k_EMsgGCStopFindingMatch:   MsgSenderClient,
	pb.EDOTAGCMsg_k_EMsgGCReadyUp:            MsgSenderClient,

	pb.EDOTAGCMsg_k_EMsgGCLeaverDetected:         MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCLeaverDetectedResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCRequestSaveGamesServer: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgDOTALiveLeagueGameUpdate: MsgSenderGC,

	pb.EDOTAGCMsg_k_EMsgGCTeamMemberProfileRequest: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCRequestPlayerResources:         MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCRequestPlayerResourcesResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCBanStatusRequest:  MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCBanStatusResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCGenerateDiretidePrizeList:         MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCGenerateDiretidePrizeListResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCPassportDataRequest:  MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCPassportDataResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCFantasyLeagueCreateInfoRequest:  MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCFantasyLeagueCreateInfoResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCFantasyLeagueInviteInfoRequest:  MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCFantasyLeagueInviteInfoResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCFantasyLeagueFriendJoinListRequest:  MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCFantasyLeagueFriendJoinListResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCRequestBatchPlayerResources:         MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCRequestBatchPlayerResourcesResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCToClientLeaguePredictionsResponse: MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgClientToGCLeaguePredictions:         MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyLeaveResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgClientToGCSetSpectatorLobbyDetailsResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgClientToGCCreateSpectatorLobbyResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptionsResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgClientEconNotification_Job: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgTeamFanfare:         MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgResponseTeamFanfare: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgDOTAAwardEventPoints: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgDOTAFrostivusTimeElapsed: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCDev_GrantWarKill: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCLobbyList:         MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCLobbyListResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCCreateFantasyLeagueRequest:  MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCCreateFantasyLeagueResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCCreateFantasyTeamRequest:  MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCCreateFantasyTeamResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCConnectedPlayers: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCLeagueAdminList: MsgSenderGC,

	pb.EDOTAGCMsg_k_EMsgGCChatMessage: MsgSenderClient,

	// Hand-written lobby code
	pb.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot: MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCPracticeLobbyCreate:      MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCFantasyLivePlayerStats: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCFantasyMatch:             MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCPCBangTimedRewardMessage: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgCastMatchVoteResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardStats:         MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardStatsResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCToClientProfileCardStatsUpdated: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCToClientAutomatedTournamentStateChange: MsgSenderNone,

	// todo: determine who sends the CMsgWeekendTourneyOpts and what the response is
	pb.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyOpts:         MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyOptsResponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPNotifyRecipient: MsgSenderGC,
	pb.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded:         MsgSenderGC,

	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC:                  MsgSenderClient,
	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter: MsgSenderGC,
	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee:           MsgSenderGC,
	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC:          MsgSenderClient,
	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInvitee:          MsgSenderClient,
	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter:          MsgSenderGC,

	pb.EDOTAGCMsg_k_EMsgGCToClientAllStarVotesReply:   MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCToClientAllStarVotesRequest: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCToClientAllStarVotesSubmit:      MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCToClientAllStarVotesSubmitReply: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgDOTALeagueInfoListAdminsRequest: MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgDOTALeagueInfoListAdminsReponse: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCtoServerTensorflowInstance: MsgSenderNone,

	pb.EDOTAGCMsg_k_EMsgGCBalancedShuffleLobby: MsgSenderClient,
	pb.EDOTAGCMsg_k_EMsgGCWatchGame:            MsgSenderClient,

	pb.EDOTAGCMsg_k_EMsgClientToGCGetUnderlordsCDKeyRequest:  MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCtoGCRequestRecalibrationCheck:      MsgSenderNone,
	pb.EDOTAGCMsg_k_EMsgGCtoGCAssociatedExploiterAccountInfo: MsgSenderNone,
}

// msgMethodNameOverrides overrides the generated client method names.
var msgMethodNameOverrides = map[pb.EDOTAGCMsg]string{
	pb.EDOTAGCMsg_k_EMsgGameAutographReward:               "AutographReward",
	pb.EDOTAGCMsg_k_EMsgDestroyLobbyRequest:               "DestroyLobby",
	pb.EDOTAGCMsg_k_EMsgGCReadyUp:                         "SendReadyUp",
	pb.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame:              "AbandonLobby",
	pb.EDOTAGCMsg_k_EMsgGCDOTAClearNotifySuccessfulReport: "ClearSuccessfulReportNotification",
	pb.EDOTAGCMsg_k_EMsgClientToGCGetTrophyList:           "ListTrophies",
	pb.EDOTAGCMsg_k_EMsgGCCreateFantasyTeamRequest:        "CreateFantasyTeam",
	pb.EDOTAGCMsg_k_EMsgGCEditFantasyTeamRequest:          "EditFantasyTeam",
	pb.EDOTAGCMsg_k_EMsgClientToGCPrivateChatKick:         "KickPrivateChatMember",
	pb.EDOTAGCMsg_k_EMsgClientToGCPrivateChatPromote:      "PromotePrivateChatMember",
	pb.EDOTAGCMsg_k_EMsgClientToGCPrivateChatDemote:       "DemotePrivateChatMember",
	pb.EDOTAGCMsg_k_EMsgClientToGCPrivateChatInfoRequest:  "RequestPrivateChatInfo",
	pb.EDOTAGCMsg_k_EMsgClientToGCPrivateChatInvite:       "InvitePrivateChatMember",
	pb.EDOTAGCMsg_k_EMsgCastMatchVote:                     "CastMatchVote",
	pb.EDOTAGCMsg_k_EMsgGCPracticeLobbyKick:               "KickLobbyMember",
	pb.EDOTAGCMsg_k_EMsgGCPracticeLobbyKickFromTeam:       "KickLobbyMemberFromTeam",
	pb.EDOTAGCMsg_k_EMsgGCBotGameCreate:                   "CreateBotGame",
	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC:          "InvitePlayerToTeam",
	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC:  "RespondToTeamInvite",
	pb.EDOTAGCMsg_k_EMsgClientsRejoinChatChannels:         "RejoinAllChatChannels",
	pb.EDOTAGCMsg_k_EMsgPartyReadyCheckRequest:            "SendPartyReadyCheck",
	pb.EDOTAGCMsg_k_EMsgPartyReadyCheckAcknowledge:        "AckPartyReadyCheck",
}

// msgResponseOverrides maps request message IDs to response message IDs.
// Setting zero as the value indicates it is an action and not a query
var msgResponseOverrides = map[pb.EDOTAGCMsg]pb.EDOTAGCMsg{
	// Example:
	// pb.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPack: pb.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPackResponse,
	pb.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest: pb.EDOTAGCMsg_k_EMsgGCToClientTeamInfo,

	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC:         pb.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter,
	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC: pb.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInvitee,
	pb.EDOTAGCMsg_k_EMsgGCWatchGame:                      pb.EDOTAGCMsg_k_EMsgGCWatchGameResponse,

	pb.EDOTAGCMsg_k_EMsgGCBalancedShuffleLobby:         0,
	pb.EDOTAGCMsg_k_EMsgGCNotificationsMarkReadRequest: 0,
}

// msgProtoTypeOverrides overrides the GC message to proto mapping.
var msgProtoTypeOverrides = map[pb.EDOTAGCMsg]proto.Message{
	pb.EDOTAGCMsg_k_EMsgGCToClientTeamInfo: &pb.CMsgDOTATeamInfo{},

	pb.EDOTAGCMsg_k_EMsgGCCreateFantasyTeamRequest:  &pb.CMsgDOTAFantasyTeamCreateRequest{},
	pb.EDOTAGCMsg_k_EMsgGCCreateFantasyTeamResponse: &pb.CMsgDOTAFantasyTeamCreateResponse{},

	pb.EDOTAGCMsg_k_EMsgGCCompendiumSetSelection:         &pb.CMsgDOTACompendiumSelection{},
	pb.EDOTAGCMsg_k_EMsgGCCompendiumSetSelectionResponse: &pb.CMsgDOTACompendiumSelectionResponse{},

	pb.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecard:        &pb.CMsgPlayerConductScorecard{},
	pb.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecardRequest: &pb.CMsgPlayerConductScorecardRequest{},

	pb.EDOTAGCMsg_k_EMsgClientToGCEventGoalsResponse: &pb.CMsgEventGoals{},

	pb.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyOptsResponse:           &pb.CMsgWeekendTourneyOpts{},
	pb.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyLeave:                  &pb.CMsgWeekendTourneyLeave{},
	pb.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStatsResponse: &pb.CMsgDOTAWeekendTourneyPlayerStats{},
	pb.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStats:         &pb.CMsgDOTAWeekendTourneyPlayerStatsRequest{},
	pb.EDOTAGCMsg_k_EMsgDOTAGetWeekendTourneySchedule:                  &pb.CMsgRequestWeekendTourneySchedule{},

	pb.EDOTAGCMsg_k_EMsgClientToGCSetPartyLeader:     &pb.CMsgDOTASetGroupLeader{},
	pb.EDOTAGCMsg_k_EMsgClientToGCCancelPartyInvites: &pb.CMsgDOTACancelGroupInvites{},

	pb.EDOTAGCMsg_k_EMsgClientToGCSetPartyOpen: &pb.CMsgDOTASetGroupOpenStatus{},

	pb.EDOTAGCMsg_k_EMsgClientToGCMergePartyInvite:        &pb.CMsgDOTAGroupMergeInvite{},
	pb.EDOTAGCMsg_k_EMsgClientToGCMergePartyResponse:      &pb.CMsgDOTAGroupMergeResponse{},
	pb.EDOTAGCMsg_k_EMsgGCToClientMergePartyResponseReply: &pb.CMsgDOTAGroupMergeReply{},
	pb.EDOTAGCMsg_k_EMsgGCToClientMergeGroupInviteReply:   &pb.CMsgDOTAGroupMergeReply{},

	pb.EDOTAGCMsg_k_EMsgClientToGCPingData: &pb.CMsgClientPingData{},

	pb.EDOTAGCMsg_k_EMsgClientToGCEventGoalsRequest: &pb.CMsgClientToGCGetEventGoals{},

	pb.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest: &pb.CMsgDOTAMyTeamInfoRequest{},

	pb.EDOTAGCMsg_k_EMsgLobbyBattleCupVictory: &pb.CMsgBattleCupVictory{},

	pb.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptions: &pb.CMsgPartyBuilderOptions{},

	pb.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel: &pb.CMsgDOTAOtherJoinedChatChannel{},
	pb.EDOTAGCMsg_k_EMsgGCOtherLeftChannel:   &pb.CMsgDOTAOtherLeftChatChannel{},

	pb.EDOTAGCMsg_k_EMsgGCCompendiumDataChanged: &pb.CMsgDOTACompendiumData{},

	pb.EDOTAGCMsg_k_EMsgGCToClientProfileCardUpdated:   &pb.CMsgDOTAProfileCard{},
	pb.EDOTAGCMsg_k_EMsgGCToClientNotificationsUpdated: &pb.CMsgGCNotificationsResponse{},

	pb.EDOTAGCMsg_k_EMsgRetrieveMatchVoteResponse: &pb.CMsgMatchVoteResponse{},

	pb.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardResponse: &pb.CMsgDOTAProfileCard{},

	pb.EDOTAGCMsg_k_EMsgGCToClientChatRegionsEnabled: &pb.CMsgDOTAChatRegionsEnabled{},

	pb.EDOTAGCMsg_k_EMsgClientToGCGetProfileTicketsResponse: &pb.CMsgDOTAProfileTickets{},

	// Experimental
	pb.EDOTAGCMsg_k_EMsgGCFantasyFinalPlayerStats: &pb.CMsgDOTAFantasyPlayerHisoricalStatsResponse_PlayerStats{},

	pb.EDOTAGCMsg_k_EMsgGCToClientTeamsInfo: &pb.CMsgDOTATeamsInfo{},

	pb.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPNotifyRecipient: &pb.CMsgDOTALobbyMVPNotifyRecipient{},
	pb.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded:         &pb.CMsgDOTALobbyMVPAwarded{},

	pb.EDOTAGCMsg_k_EMsgClientToGCRequestEventTipsSummary:         &pb.CMsgEventTipsSummaryRequest{},
	pb.EDOTAGCMsg_k_EMsgClientToGCRequestEventTipsSummaryResponse: &pb.CMsgEventTipsSummaryResponse{},

	pb.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeed:         &pb.CMsgSocialFeedRequest{},
	pb.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedResponse: &pb.CMsgSocialFeedResponse{},

	pb.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedComments:         &pb.CMsgSocialFeedCommentsRequest{},
	pb.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedCommentsResponse: &pb.CMsgSocialFeedCommentsResponse{},
}

var msgArgAsParameterOverrides = map[pb.EDOTAGCMsg]bool{
	pb.EDOTAGCMsg_k_EMsgGCPracticeLobbySetDetails: true,
}

var msgEventNameOverrides = map[pb.EDOTAGCMsg]string{
	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee:  "TeamInviteReceived",
	pb.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter: "TeamInviteResponseReceived",
	pb.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel:             "PlayerJoinedChannel",
	pb.EDOTAGCMsg_k_EMsgGCOtherLeftChannel:               "PlayerLeftChannel",
}
