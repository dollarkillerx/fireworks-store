package utils

import (
	cfg "github.com/dollarkillerx/common/pkg/config"
	"github.com/dollarkillerx/common/pkg/logger"
	"github.com/dollarkillerx/creeper/sdk/creeper_sdk"

	"log"
)

var Logger *logger.RimeLogger

func InitLogger(config cfg.LoggerConfig, creeperConfig cfg.CreeperConfiguration) {
	sdk := creeper_sdk.New(creeperConfig.Uri, creeperConfig.Token)
	_, err := sdk.Index()
	if err != nil {
		panic(err)
	}
	Logger = logger.NewLogger().
		Level(config.Level.Level()).
		Formatter(logger.DefaultFormatter()).
		Rotation(logger.DefaultCustomizeLog(&config, func(p []byte) {
			err := sdk.Log("fireworks", string(p))
			if err != nil {
				log.Println(err)
			}
		})).
		SetLogReportCaller(true).
		Complete()

}
