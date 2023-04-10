package nex

import (
	"github.com/PretendoNetwork/mario-kart-8-secure/globals"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/match-making-ext"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/nat-traversal"
	"github.com/PretendoNetwork/nex-protocols-go/ranking"
	secure_connection "github.com/PretendoNetwork/nex-protocols-go/secure-connection"

	nex_match_making "github.com/PretendoNetwork/mario-kart-8-secure/nex/match-making"
	nex_match_making_ext "github.com/PretendoNetwork/mario-kart-8-secure/nex/match-making-ext"
	nex_matchmake_extension "github.com/PretendoNetwork/mario-kart-8-secure/nex/matchmake-extension"
	nex_nat_traversal "github.com/PretendoNetwork/mario-kart-8-secure/nex/nat-traversal"
	nex_ranking "github.com/PretendoNetwork/mario-kart-8-secure/nex/ranking"
	nex_secure_connection "github.com/PretendoNetwork/mario-kart-8-secure/nex/secure-connection"
)

func registerNEXProtocols() {
	secureConnectionProtocol := secure_connection.NewSecureConnectionProtocol(globals.NEXServer)

	secureConnectionProtocol.Register(nex_secure_connection.Register)
	secureConnectionProtocol.ReplaceURL(nex_secure_connection.ReplaceURL)
	secureConnectionProtocol.SendReport(nex_secure_connection.SendReport)

	natTraversalProtocol := nat_traversal.NewNATTraversalProtocol(globals.NEXServer)

	natTraversalProtocol.RequestProbeInitiationExt(nex_nat_traversal.RequestProbeInitiationExt)
	natTraversalProtocol.ReportNATProperties(nex_nat_traversal.ReportNATProperties)

	matchmakeExtensionProtocol := matchmake_extension.NewMatchmakeExtensionProtocol(globals.NEXServer)

	matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_Postpone(nex_matchmake_extension.AutoMatchmakeWithSearchCriteria_Postpone)

	matchMakingProtocol := match_making.NewMatchMakingProtocol(globals.NEXServer)

	matchMakingProtocol.GetSessionURLs(nex_match_making.GetSessionURLs)
	matchMakingProtocol.UpdateSessionHostV1(nex_match_making.UpdateSessionHostV1)

	matchMakingExtProtocol := match_making_ext.NewMatchMakingExtProtocol(globals.NEXServer)

	matchMakingExtProtocol.EndParticipation(nex_match_making_ext.EndParticipation)

	rankingProtocol := ranking.NewRankingProtocol(globals.NEXServer)

	rankingProtocol.UploadCommonData(nex_ranking.UploadCommonData)
}
