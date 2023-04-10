package nex_secure_connection

import (
	"strconv"

	"github.com/PretendoNetwork/mario-kart-8-secure/database"
	"github.com/PretendoNetwork/mario-kart-8-secure/globals"
	nex "github.com/PretendoNetwork/nex-go"
	secure_connection "github.com/PretendoNetwork/nex-protocols-go/secure-connection"
)

func Register(err error, client *nex.Client, callID uint32, stationUrls []*nex.StationURL) {
	localStation := stationUrls[0]
	localStationURL := localStation.EncodeToString()
	connectionId := uint32(globals.NEXServer.ConnectionIDCounter().Increment())
	client.SetConnectionID(connectionId)
	client.SetLocalStationURL(localStationURL)

	address := client.Address().IP.String()
	port := strconv.Itoa(client.Address().Port)
	natf := "0"
	natm := "0"
	type_ := "3"

	localStation.SetAddress(address)
	localStation.SetPort(port)
	localStation.SetNatf(natf)
	localStation.SetNatm(natm)
	localStation.SetType(type_)

	globalStationURL := localStation.EncodeToString()

	if !database.DoesSessionExist(client.PID()) {
		database.AddPlayerSession(client.PID(), []string{localStationURL, globalStationURL}, address, port)
	} else {
		database.UpdatePlayerSessionAll(client.PID(), []string{localStationURL, globalStationURL}, address, port)
	}

	rmcResponseStream := nex.NewStreamOut(globals.NEXServer)

	rmcResponseStream.WriteUInt32LE(0x10001) // Success
	rmcResponseStream.WriteUInt32LE(connectionId)
	rmcResponseStream.WriteString(globalStationURL)

	rmcResponseBody := rmcResponseStream.Bytes()

	// Build response packet
	rmcResponse := nex.NewRMCResponse(secure_connection.ProtocolID, callID)
	rmcResponse.SetSuccess(secure_connection.MethodRegister, rmcResponseBody)

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
