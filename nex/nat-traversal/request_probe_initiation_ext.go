package nex_nat_traversal

import (
	"strconv"

	"github.com/PretendoNetwork/mario-kart-8-secure/globals"
	nex "github.com/PretendoNetwork/nex-go"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/nat-traversal"
)

func RequestProbeInitiationExt(err error, client *nex.Client, callID uint32, targetList []string, stationToProbe string) {
	rmcResponse := nex.NewRMCResponse(nat_traversal.ProtocolID, callID)
	rmcResponse.SetSuccess(nat_traversal.MethodRequestProbeInitiationExt, nil)

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

	rmcMessage := nex.RMCRequest{}
	rmcMessage.SetProtocolID(nat_traversal.ProtocolID)
	rmcMessage.SetCallID(0xffff0000 + callID)
	rmcMessage.SetMethodID(nat_traversal.MethodInitiateProbe)
	rmcRequestStream := nex.NewStreamOut(globals.NEXServer)
	rmcRequestStream.WriteString(stationToProbe)
	rmcRequestBody := rmcRequestStream.Bytes()
	rmcMessage.SetParameters(rmcRequestBody)
	rmcMessageBytes := rmcMessage.Bytes()

	for _, target := range targetList {
		targetUrl := nex.NewStationURL(target)
		targetRvcID, _ := strconv.Atoi(targetUrl.RVCID())
		targetClient := globals.NEXServer.FindClientFromConnectionID(uint32(targetRvcID))
		if targetClient != nil {
			messagePacket, _ := nex.NewPacketV1(targetClient, nil)
			messagePacket.SetVersion(1)
			messagePacket.SetSource(0xA1)
			messagePacket.SetDestination(0xAF)
			messagePacket.SetType(nex.DataPacket)
			messagePacket.SetPayload(rmcMessageBytes)

			messagePacket.AddFlag(nex.FlagNeedsAck)
			messagePacket.AddFlag(nex.FlagReliable)

			globals.NEXServer.Send(messagePacket)
		}
	}
}
