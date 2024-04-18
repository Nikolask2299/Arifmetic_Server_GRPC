.DEFAULT_GLOBAL := gen

building:
	go build -o arifm-server services/cmd/server/main.go 
.PHONY:building

running:
	go run services/cmd/migrations/main.go --storage-path=./storage/storage.db --migration-path=./migration
.PHONY:running