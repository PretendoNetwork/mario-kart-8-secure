package nex_match_making_ext

import (
	"github.com/PretendoNetwork/mario-kart-8-secure/database"
	"github.com/PretendoNetwork/mario-kart-8-secure/globals"
	nex "github.com/PretendoNetwork/nex-go"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/match-making-ext"
)

func EndParticipation(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string) {
	database.RemovePlayerFromRoom(idGathering, client.PID())

	returnval := []byte{0x1}

	rmcResponse := nex.NewRMCResponse(match_making_ext.ProtocolID, callID)
	rmcResponse.SetSuccess(match_making_ext.MethodEndParticipation, returnval)

	rmcResponseBytes := rmcResponse.Bytes()

	responsePacket, _ := nex.NewPacketV1(client, nil)

	responsePacket.SetVersion(1)
	responsePacket.SetSource(0xA1)
	responsePacket.SetDestination(0xAF)
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	globals.NEXServer.Send(responsePacket)
}
