package config

const(
	// Service ports
	ItemSaverPort = 1234

	// elasticsearch
	ElasticIndex = "dating_profile"
	ElasticUrl = "http://192.168.99.100:9200"

	// RPC endpoints
	ItemSaverRpc = "ItemSaverService.Save"

	// Local server ip address
	ServerIpDefault = "127.0.0.1" // localhost
	ServerIpVirtualBox = "192.168.99.100" // virtualbox ip
)
