package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dollarkillerx/fireworks/internal/config"
	"github.com/dollarkillerx/fireworks/internal/generated"
	"github.com/dollarkillerx/fireworks/internal/middlewares"
	"github.com/dollarkillerx/fireworks/internal/resolvers"
	"github.com/dollarkillerx/fireworks/internal/storage/basis"
	"github.com/dollarkillerx/fireworks/pkg/utils"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
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
	utils.InitLogger(config.GetLoggerConfig(), config.GetCreeperConfig())
	postgres, err := utils.InitPostgres(config.GetPostgresConfig())
	if err != nil {
		log.Fatalln(err)
	}

	storage, err := basis.New(postgres)
	if err != nil {
		log.Fatalln("Basis Storage 初始化失败: ", err)
	}

	// 初始化中间件 End

	router := chi.NewRouter()

	router.Use(middlewares.Cors())
	router.Use(middlewares.Context())

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ack"))
	})

	if config.GetEnablePlayground() {
		router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
		utils.Logger.Warningln("playground is enabled!")
	}

	conf := generated.Config{
		Resolvers: resolvers.NewResolver(storage),
	}

	conf.Directives.HasLogined = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		//userID, err := middlewares.GetUserIDFromCtx(ctx)
		//if err != nil {
		//	return nil, err
		//}
		//ctx = middlewares.SetContextUserID(ctx, userID)
		return next(ctx)
	}

	graphQLServer := newServer(generated.NewExecutableSchema(conf))
	graphQLServer.SetErrorPresenter(middlewares.ErrorPresenter)
	graphQLServer.SetRecoverFunc(middlewares.RecoverFunc)
	graphQLServer.Use(&middlewares.RequestLogger{})
	graphQLServer.AroundResponses(middlewares.ResponseLogger)
	router.Handle("/graphql", graphQLServer)

	utils.Logger.Infof("connect to http://localhost:%s/playground for GraphQL playground", getPort())

	if err := http.ListenAndServe(":"+getPort(), router); err != nil {
		panic(err)
	}
}

const InstancePort = "INSTANCE_PORT"

func getPort() string {
	port := os.Getenv(InstancePort)
	if port == "" {
		return "8181"
	}

	return port
}

func newServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			Subprotocols:    []string{"graphql-transport-ws", "graphql-ws"},
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return srv
}
