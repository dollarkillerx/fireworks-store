package basis

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/dollarkillerx/fireworks/internal/config"
	"github.com/dollarkillerx/fireworks/pkg/utils"
)

var configFilename = "config"
var configDirs = "./,../../../configs/"
var storage *Storage

func init() {
	err := config.InitConfig(configFilename, strings.Split(configDirs, ","))
	if err != nil {
		panic(fmt.Errorf("error parsing config, %s", err))
	}

	postgres, err := utils.InitPostgres(config.GetPostgresConfig())
	if err != nil {
		log.Fatalln(err)
	}
	storage, err = New(postgres)
	if err != nil {
		log.Fatalln("Basis Storage 初始化失败: ", err)
	}

}

func TestGetShopList(t *testing.T) {
	total, shops, err := storage.GetShopList(true, 0, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("商铺数量：", total)
	fmt.Println("商铺信息：", shops)

}
