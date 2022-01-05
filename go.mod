module github.com/dollarkillerx/fireworks

go 1.15

require (
	github.com/99designs/gqlgen v0.14.1-0.20211213161003-27a2b210d913
	github.com/dollarkillerx/common v0.0.2
	github.com/vektah/gqlparser/v2 v2.2.0
)

require (
	github.com/dollarkillerx/urllib v1.13.16
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.6.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/postgres v1.2.3 // indirect
	gorm.io/gorm v1.22.4
)

replace (
	github.com/dollarkillerx/common v0.0.2 => /home/wangy/workspace/2022/common
)
