dev:
	@echo "Starting dev mode"
	@npx tailwind -i ./style/main.css -o ./public/main.css --watch & \
		air

serve:
	@npx tailwind -i ./style/main.css -o ./public/main.css & \
		go run cmd/main.go -dev -port 3000

build:
	@echo "Building for prod"
	@go run cmd/main.go
	@npx tailwind -i ./style/main.css -o ./dist/main.css --minify
