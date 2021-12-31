graphql:
	go get -d github.com/99designs/gqlgen@v0.14.1-0.20211213161003-27a2b210d913
	rm -f ./internal/generated/resolver.go
	go run github.com/99designs/gqlgen generate --config ./configs/gqlgen.yaml --verbose
