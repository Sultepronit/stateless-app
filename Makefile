.PHONY: build deploy

SERVICE=go-lang-service

build: 
	@echo "Builging..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/app .
	@echo "Done!"

deploy: build
	@echo "Deploying..."
	cd ../../.. && docker compose restart $(SERVICE)
	@echo "Done!"