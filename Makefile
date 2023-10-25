buildapi :
	@go build -o ./bin/cmd/api/api ./cmd/api/api.go

runapi: buildapi
	@./bin/cmd/api/api
