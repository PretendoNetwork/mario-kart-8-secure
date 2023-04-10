package main

import (
	"github.com/PretendoNetwork/mario-kart-8-secure/database"
	"github.com/PretendoNetwork/mario-kart-8-secure/globals"
	"github.com/PretendoNetwork/mario-kart-8-secure/utility"
)

func init() {

	globals.Config, _ = utility.ImportConfigFromFile("secure.config")

	database.ConnectAll()

	globals.MatchmakingState = append(globals.MatchmakingState, nil)
}
