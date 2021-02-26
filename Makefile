setup-dev:
	GO111MODULE=off go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
	GO111MODULE=off go get github.com/google/wire/cmd/wire
	GO111MODULE=off go get -u github.com/cosmtrek/air
	GO111MODULE=off go get -u github.com/swaggo/swag/cmd/swag

start-dev:
	docker-compose --project-name mcm -f ./deployments/docker-compose-dev.yaml up -d
	air -c ./scripts/air.toml

build-air:
	swag init -g internal/server/server.go
	go build -o ./tmp/main .