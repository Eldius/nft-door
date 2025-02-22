package configs

import "github.com/spf13/viper"

const (
	AppName = "nft-pocs"
)

func GetNetworkEndpoint() string {
	return viper.GetString("eth.network.endpoint")
}

func GetPrivateKey() string {
	return viper.GetString("eth.network.private_key")
}
