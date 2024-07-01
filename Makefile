dev:
	@echo "Starting dev mode"
	APP_ENV="pre" air & \
		npx tailwind -i ./style/main.css -o ./dist/main.css --watch

build:
	@echo "Building for prod"
	APP_ENV="pro" go run cmd/main.go
	@npx tailwind -i ./style/main.css -o ./dist/main.css --minify
