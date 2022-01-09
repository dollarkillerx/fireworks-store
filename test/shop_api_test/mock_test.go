package shop_api_test

import (
	"fmt"
	"github.com/dollarkillerx/fireworks/internal/config"
	"github.com/dollarkillerx/fireworks/pkg/models"
	"github.com/dollarkillerx/fireworks/pkg/utils"
	"github.com/dollarkillerx/warehouse/pkg/warehouse_sdk"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

var configFilename string
var configDirs string

func init() {
	//const (
	//	defaultConfigFilename = "config"
	//	configUsage           = "Name of the config file, without extension"
	//	defaultConfigDirs     = "./,../../configs/"
	//	configDirUsage        = "Directories to search for config file, separated by ','"
	//)
	//flag.StringVar(&configFilename, "c", defaultConfigFilename, configUsage)
	//flag.StringVar(&configDirs, "cPath", defaultConfigDirs, configDirUsage)
	//flag.Parse()
	//fmt.Println("init success")
	configFilename = "config"
	configDirs = "./,../../configs/"
}

var lgo = "/home/cg/Pictures/34052213.png"
var img = "/home/cg/Pictures/chengxiao.jpeg"

func TestUploadImage(t *testing.T) {
	sdk := warehouse_sdk.New("http://192.168.31.20:8187", "access_key", "secret_key", 0)
	err := sdk.Ping()
	if err != nil {
		panic(err)
	}

	file, err := ioutil.ReadFile(lgo)
	if err != nil {
		panic(err)
	}
	err = sdk.PutObject("img", lgo, file, nil)
	if err != nil {
		panic(err)
	}

	object, err := sdk.GetObject("img", lgo)
	if err != nil {
		panic(err)
	}

	fmt.Println(object.MetaData)
	fmt.Println(len(object.Data))

	//err = sdk.DelObject("img", img)
	//if err != nil {
	//	panic(err)
	//}

	//err = sdk.RemoveBucket("img")
	//if err != nil {
	//	panic(err)
	//}
}

func TestAddShop(t *testing.T) {

	err := config.InitConfig(configFilename, strings.Split(configDirs, ","))
	if err != nil {
		panic(fmt.Errorf("error parsing config, %s", err))
	}

	postgres, err := utils.InitPostgres(config.GetPostgresConfig())
	if err != nil {
		log.Fatalln(err)
	}

	postgres.Create(&models.Shop{
		CommonModel:   models.CommonModel{ID: utils.SonuFlakeStringID()},
		Name:          "小王商铺",
		Phone:         "15690890221",
		LogoUri:       "http://192.168.31.20:8187/v1/download?file=/home/cg/Pictures/34052213.png&bucket=img",
		Address:       "成都市双流区吉龙二创业花园150号",
		Weights:       8,
		Announcement:  "开业大吉",
		Score:         4.9,
		StartDelivery: 10,
		GPSLatitude:   0,
		GPSLongitude:  0,
		Account:       "admin@xw",
		Password:      "123456",
		Activation:    true,
		Admin:         true,
	})

}
