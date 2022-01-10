package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/dollarkillerx/fireworks/internal/config"
	"github.com/dollarkillerx/fireworks/pkg/models"
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

	for i := 0; i < 10; i++ {
		err := postgres.Model(&models.Shop{}).Create(&models.Shop{
			CommonModel:   models.DefaultCommonModel(),
			Name:          fmt.Sprintf("商铺: %d", i),
			Phone:         "189528514",
			LogoUri:       "?file=f5e0a144c95fcf4d0d9d9b9770a51f6d39970.jpg&bucket=img",
			Weights:       float64(rand.Intn(5)),
			Announcement:  "全场大放送",
			Score:         float64(rand.Intn(5)),
			StartDelivery: 20,
			Account:       fmt.Sprintf("qq_%d", i),
			Password:      utils.Md5Encode("123456" + config.GetJWTConfig().Salt),
			Activation:    true,
			Admin:         true,
		}).Error
		if err != nil {
			panic(err)
		}
	}
}
