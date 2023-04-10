package nex_secure_connection

import (
	"github.com/PretendoNetwork/mario-kart-8-secure/database"
	"github.com/PretendoNetwork/mario-kart-8-secure/globals"
	nex "github.com/PretendoNetwork/nex-go"
	secure_connection "github.com/PretendoNetwork/nex-protocols-go/secure-connection"
)

func ReplaceURL(err error, client *nex.Client, callID uint32, oldStation *nex.StationURL, newStation *nex.StationURL) {
	database.UpdatePlayerSessionURL(client.PID(), oldStation.EncodeToString(), newStation.EncodeToString())

	rmcResponse := nex.NewRMCResponse(secure_connection.ProtocolID, callID)
	rmcResponse.SetSuccess(secure_connection.MethodReplaceURL, nil)

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
