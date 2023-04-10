package nex_matchmake_extension

import (
	"fmt"
	"math"
	"strconv"

	"github.com/PretendoNetwork/mario-kart-8-secure/database"
	"github.com/PretendoNetwork/mario-kart-8-secure/globals"
	nex "github.com/PretendoNetwork/nex-go"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
)

var regionList = []string{"Worldwide", "Japan", "United States", "Europe", "Korea", "China", "Taiwan"}
var gameModes = []string{"VS Race", "Battle"}
var ccList = []string{"Unk", "200cc", "50cc", "100cc", "150cc", "Mirror", "BattleCC"}
var itemModes = []string{"Unk1", "Unk2", "Unk3", "Unk4", "Unk5", "Normal", "Unk7", "All Items", "Shells Only", "Bananas Only", "Mushrooms Only", "Bob-ombs Only", "No Items", "No Items or Coins", "Frantic"}
var vehicleModes = []string{"All Vehicles", "Karts Only", "Bikes Only"}
var controllerModes = []string{"Unk", "Tilt Only", "All Controls"}
var dlcModes = []string{"No DLC", "DLC Pack 1 Only", "DLC Pack 2 Only", "Both DLC Packs"}

func AutoMatchmakeWithSearchCriteria_Postpone(err error, client *nex.Client, callID uint32, searchCriteria []*match_making.MatchmakeSessionSearchCriteria, anyGathering *nex.DataHolder, message string) {

	// * From Jon: This was added here because this function had the wrong signature
	// * I have no idea if this works at all, I just got it to build
	matchmakeSession := anyGathering.ObjectData().(*match_making.MatchmakeSession)

	gameConfig := matchmakeSession.Attributes[2]
	fmt.Println(strconv.FormatUint(uint64(gameConfig), 2))
	fmt.Println("===== MATCHMAKE SESSION JOIN =====")
	fmt.Println("REGION: " + regionList[matchmakeSession.Attributes[3]])
	fmt.Println("GAME MODE: " + gameModes[matchmakeSession.GameMode])
	fmt.Println("CC: " + ccList[gameConfig%0b111])
	gameConfig = gameConfig >> 12
	fmt.Println("DLC MODE: " + dlcModes[matchmakeSession.Attributes[5]&0xF])
	fmt.Println("ITEM MODE: " + itemModes[gameConfig%0b1111])
	gameConfig = gameConfig >> 8
	fmt.Println("VEHICLE MODE: " + vehicleModes[gameConfig%0b11])
	gameConfig = gameConfig >> 4
	fmt.Println("CONTROLLER MODE: " + controllerModes[gameConfig%0b11])
	fmt.Println("HAVE GUEST PLAYER: " + strconv.FormatBool(searchCriteria[0].VacantParticipants > 1))

	gid := database.FindRoom(matchmakeSession.GameMode, true, matchmakeSession.Attributes[3], matchmakeSession.Attributes[2], uint32(searchCriteria[0].VacantParticipants), matchmakeSession.Attributes[5]&0xF)
	if gid == math.MaxUint32 {
		gid = database.NewRoom(client.PID(), matchmakeSession.GameMode, true, matchmakeSession.Attributes[3], matchmakeSession.Attributes[2], uint32(searchCriteria[0].VacantParticipants), matchmakeSession.Attributes[5]&0xF)
	}

	database.AddPlayerToRoom(gid, client.PID(), uint32(searchCriteria[0].VacantParticipants))

	hostpid, gamemode, region, gconfig, dlcmode := database.GetRoomInfo(gid)
	sessionKey := "00000000000000000000000000000000"

	matchmakeSession.Gathering.ID = gid
	matchmakeSession.Gathering.OwnerPID = hostpid
	matchmakeSession.Gathering.HostPID = hostpid
	matchmakeSession.Gathering.MinimumParticipants = 1
	matchmakeSession.SessionKey = []byte(sessionKey)
	matchmakeSession.GameMode = gamemode
	matchmakeSession.Attributes[3] = region
	matchmakeSession.Attributes[2] = gconfig
	matchmakeSession.Attributes[5] = dlcmode

	rmcResponseStream := nex.NewStreamOut(globals.NEXServer)
	rmcResponseStream.WriteString("MatchmakeSession")
	lengthStream := nex.NewStreamOut(globals.NEXServer)
	lengthStream.WriteStructure(matchmakeSession.Gathering)
	lengthStream.WriteStructure(matchmakeSession)
	matchmakeSessionLength := uint32(len(lengthStream.Bytes()))
	rmcResponseStream.WriteUInt32LE(matchmakeSessionLength + 4)
	rmcResponseStream.WriteUInt32LE(matchmakeSessionLength)
	rmcResponseStream.WriteStructure(matchmakeSession.Gathering)
	rmcResponseStream.WriteStructure(matchmakeSession)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCResponse(matchmake_extension.ProtocolID, callID)
	rmcResponse.SetSuccess(matchmake_extension.MethodAutoMatchmakeWithSearchCriteria_Postpone, rmcResponseBody)

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
