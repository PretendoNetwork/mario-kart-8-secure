package types

type ServerConfig struct {
	ServerName            string
	ServerPort            string
	PrudpVersion          int
	SignatureVersion      int
	KerberosKeySize       int
	AccessKey             string
	NexVersion            int
	DatabaseIP            string
	DatabasePort          string
	DatabaseUseAuth       bool
	DatabaseUsername      string
	DatabasePassword      string
	AccountDatabase       string
	PNIDCollection        string
	NexAccountsCollection string
	MK8Database           string
	RoomsCollection       string
	SessionsCollection    string
	UsersCollection       string
	RegionsCollection     string
	TournamentsCollection string
}
