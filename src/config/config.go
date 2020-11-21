/*
  All Configs, Envs Here
*/
package config

import (
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type configCenter struct {
	Endpoint    string
	NamespaceId string
	AccessKey   string
	SecretKey   string
	DataId      string
	Group       string
}

type database struct {
	DbType   string
	DbName   string
	Server   string
	UserName string
	Password string
	Port     int
}

// ConfigCenter is Config Center configuations
type ConfigCenter struct {
	Acm configCenter
}

type envs struct {
	Env string

	AutoMigrate bool
	AutoSeed    bool
}

// Config is App Configs
type Config struct {
	Database database

	Environment envs
}

var sConfigCenter = ConfigCenter{}
var sConfig = Config{}
var once sync.Once

func loadAppDir() {
	// Work dir config
	if appDir := os.Getenv("APP_DIR"); "" != appDir {
		if err := os.Chdir(appDir); err != nil {
			logger().Fatalf("Set Current Dir '%v' failed\n", appDir)
		}
		appDir, err := os.Getwd()
		if err != nil {
			logger().Fatalf("Load Current Dir failed\n")
		}
		logger().Printf("APP_DIR='%v'", appDir)
	}
}

// Settings is All Configs Setting
func Settings() *Config {
	once.Do(func() {
		loadAppDir()
		if err := loadConfigCenter(); err != nil {
			logger().Printf("Load Config Center failed: %v\n", err.Error())
		} else {
			logger().Printf("Load config center success\n")
			return
		}
		if err := loadLocalConfig(); err != nil {
			logger().Fatalf("Load Local Config failed: %v\n", err.Error())
		} else {
			logger().Printf("Load config file success\n")
			return
		}
	})
	return &sConfig
}

func loadConfigCenter() (err error) {
	configCenterFile := os.Getenv("CONFIG_CENTER_FILE")
	if "" == configCenterFile {
		configCenterFile = "config.center.toml"
	}
	logger().Printf("CONFIG_CENTER_FILE='%v'", configCenterFile)
	_, err = toml.DecodeFile(configCenterFile, &sConfigCenter)
	if err != nil {
		return
	}
	acm := &sConfigCenter.Acm

	clientConfig := constant.ClientConfig{
		LogDir:         "./log/nacos",
		Endpoint:       acm.Endpoint,
		NamespaceId:    acm.NamespaceId,
		AccessKey:      acm.AccessKey,
		SecretKey:      acm.SecretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}
	logger().Printf("Connecting Config Center...\n")

	configClient, err := clients.CreateConfigClient(map[string]interface{}{"clientConfig": clientConfig})
	if err != nil {
		return
	}
	logger().Println("Loading from Config Center...")

	configData, err := configClient.GetConfig(vo.ConfigParam{DataId: acm.DataId, Group: acm.Group})
	if err != nil {
		return
	}

	_, err = toml.Decode(string(configData), &sConfig)
	return
}

func loadLocalConfig() (err error) {
	configFile := os.Getenv("CONFIG_FILE")
	if "" == configFile {
		configFile = "config.toml"
	}
	logger().Printf("CONFIG_FILE='%v'", configFile)

	_, err = toml.DecodeFile(configFile, &sConfig)
	return
}
