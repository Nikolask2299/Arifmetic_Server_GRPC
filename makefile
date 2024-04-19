.DEFAULT_GLOBAL := gen

running:
	go run services/cmd/server/main.go --config=config_file/configuration.json 
.PHONY:running

building:
	go build -o arifm-server services/cmd/server/main.go 
.PHONY:building

migration:
	go run services/cmd/migrations/main.go --storage-path=./storage/storage.db --migration-path=./migration
.PHONY:migration