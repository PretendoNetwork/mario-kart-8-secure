package prudp

import (
	"github.com/PretendoNetwork/mario-kart-8-secure/database"
	"github.com/PretendoNetwork/mario-kart-8-secure/globals"
	nex "github.com/PretendoNetwork/nex-go"
)

func Connect(packet *nex.PacketV1) {
	payload := packet.Payload()

	stream := nex.NewStreamIn(payload, globals.NEXServer)

	_, _ = stream.ReadBuffer()
	checkData, _ := stream.ReadBuffer()

	sessionKey := make([]byte, globals.NEXServer.KerberosKeySize())

	kerberos := nex.NewKerberosEncryption(sessionKey)

	checkDataDecrypted := kerberos.Decrypt(checkData)
	checkDataStream := nex.NewStreamIn(checkDataDecrypted, globals.NEXServer)

	userPID := checkDataStream.ReadUInt32LE() // User PID
	packet.Sender().SetPID(userPID)
	_ = checkDataStream.ReadUInt32LE() //CID of secure server station url
	responseCheck := checkDataStream.ReadUInt32LE()

	responseValueStream := nex.NewStreamOut(globals.NEXServer)
	responseValueStream.WriteUInt32LE(responseCheck + 1)

	responseValueBufferStream := nex.NewStreamOut(globals.NEXServer)
	responseValueBufferStream.WriteBuffer(responseValueStream.Bytes())

	globals.NEXServer.AcknowledgePacket(packet, responseValueBufferStream.Bytes())

	packet.Sender().UpdateRC4Key(sessionKey)
	packet.Sender().SetSessionKey(sessionKey)

	if !database.DoesUserExist(userPID) {
		database.AddNewUser(userPID)
	}
}
