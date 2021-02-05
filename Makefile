setup-dev:
	GO111MODULE=off go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

start-dev:
	docker-compose --project-name mcm -f ./deployments/docker-compose-dev.yaml up -d