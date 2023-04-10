package types

import (
	"github.com/PretendoNetwork/nex-go"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
)

type MatchmakingData struct {
	MatchmakeSession *match_making.MatchmakeSession
	Clients          []*nex.Client
}
