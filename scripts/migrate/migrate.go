package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/dollarkillerx/fireworks/internal/config"
	"github.com/dollarkillerx/fireworks/internal/storage/basis"
	"github.com/dollarkillerx/fireworks/pkg/utils"
)

var configFilename string
var configDirs string

func init() {
	const (
		defaultConfigFilename = "config"
		configUsage           = "Name of the config file, without extension"
		defaultConfigDirs     = "./,./configs/"
		configDirUsage        = "Directories to search for config file, separated by ','"
	)
	flag.StringVar(&configFilename, "c", defaultConfigFilename, configUsage)
	flag.StringVar(&configDirs, "cPath", defaultConfigDirs, configDirUsage)
	flag.Parse()
}

func main() {
	err := config.InitConfig(configFilename, strings.Split(configDirs, ","))
	if err != nil {
		panic(fmt.Errorf("error parsing config, %s", err))
	}

	// 初始化中间件
	postgres, err := utils.InitPostgres(config.GetPostgresConfig())
	if err != nil {
		log.Fatalln(err)
	}

	storage, err := basis.New(postgres)
	if err != nil {
		log.Fatalln("Basis Storage 初始化失败: ", err)
	}
	storage = storage
}
