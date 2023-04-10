package globals

import (
	"github.com/PretendoNetwork/mario-kart-8-secure/types"
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/plogger-go"
)

var Logger = plogger.NewLogger()
var NEXServer *nex.Server
var Config *types.ServerConfig
var MatchmakingState []*types.MatchmakingData
