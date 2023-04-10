package utility

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/PretendoNetwork/mario-kart-8-secure/types"
)

func ImportConfigFromFile(path string) (*types.ServerConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	indexes := strings.Split(string(data), "\n")
	config := &types.ServerConfig{
		ServerName:      "server",
		KerberosKeySize: 32,
	}
	for i := 0; i < len(indexes); i++ {
		index := strings.Split(indexes[i], "=")
		if len(index) != 2 {
			continue
		}
		switch index[0] {
		case "ServerName":
			config.ServerName = index[1]
			break
		case "ServerPort":
			config.ServerPort = index[1]
			break
		case "PrudpVersion":
			config.PrudpVersion, err = strconv.Atoi(index[1])
			if err != nil {
				return nil, err
			}
			break
		case "SignatureVersion":
			config.SignatureVersion, err = strconv.Atoi(index[1])
			if err != nil {
				return nil, err
			}
			break
		case "KerberosKeySize":
			config.KerberosKeySize, err = strconv.Atoi(index[1])
			if err != nil {
				return nil, err
			}
			break
		case "AccessKey":
			config.AccessKey = index[1]
			break
		case "NexVersion":
			config.NexVersion, err = strconv.Atoi(index[1])
			if err != nil {
				return nil, err
			}
			break
		case "DatabaseIP":
			config.DatabaseIP = index[1]
			break
		case "DatabasePort":
			config.DatabasePort = index[1]
			break
		case "DatabaseUseAuth":
			config.DatabaseUseAuth = (index[1] == "true")
			break
		case "DatabaseUsername":
			config.DatabaseUsername = index[1]
			break
		case "DatabasePassword":
			config.DatabasePassword = index[1]
			break
		case "AccountDatabase":
			config.AccountDatabase = index[1]
			break
		case "PNIDCollection":
			config.PNIDCollection = index[1]
			break
		case "NexAccountsCollection":
			config.NexAccountsCollection = index[1]
			break
		case "MK8Database":
			config.MK8Database = index[1]
			break
		case "RoomsCollection":
			config.RoomsCollection = index[1]
			break
		case "SessionsCollection":
			config.SessionsCollection = index[1]
			break
		case "UsersCollection":
			config.UsersCollection = index[1]
			break
		case "RegionsCollection":
			config.RegionsCollection = index[1]
			break
		case "TournamentsCollection":
			config.TournamentsCollection = index[1]
			break
		default:
			break
		}
	}
	return config, nil
}
