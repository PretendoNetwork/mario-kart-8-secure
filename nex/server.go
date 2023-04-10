package nex

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/mario-kart-8-secure/globals"
	"github.com/PretendoNetwork/mario-kart-8-secure/prudp"
	"github.com/PretendoNetwork/nex-go"
)

func StartNEXServer() {
	globals.NEXServer = nex.NewServer()
	globals.NEXServer.SetPRUDPVersion(1)
	globals.NEXServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 3,
		Minor: 5,
		Patch: 4,
	})
	globals.NEXServer.SetKerberosPassword(os.Getenv("KERBEROS_PASSWORD"))
	globals.NEXServer.SetAccessKey(globals.Config.AccessKey)

	globals.NEXServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==MK8 - Secure==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("=================")
	})

	globals.NEXServer.On("Connect", prudp.Connect)

	registerNEXProtocols()

	globals.NEXServer.Listen(":" + globals.Config.ServerPort)
}
