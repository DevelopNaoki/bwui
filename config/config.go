package config

import (
	"gopkg.in/go-ini/ini.v1"
)

type ServerConfig struct {
	Port              int
	SSL               bool
	ServerCertificate string
	PrivateKey        string
}

func GetServerConfig() (serverConfig ServerConfig, err error) {
	// Read Config file
	configFile, err := ini.Load("bwui.conf")
	if err != nil {
		return serverConfig, err
	}

	// Check Configuration from Config File
	serverConfig.Port = configFile.Section("server").Key("port").MustInt(8080)
	serverConfig.SSL = configFile.Section("server").Key("SSL").MustBool(false)
	if serverConfig.SSL {
		serverConfig.ServerCertificate = configFile.Section("server").Key("server_certificate").String()
		serverConfig.PrivateKey = configFile.Section("server").Key("private_key").String()
	}
	return serverConfig, nil
}
