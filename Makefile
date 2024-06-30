dev:
	@echo "Starting dev mode"
	APP_ENV="dev" air & \
		npx tailwind -i ./style/main.css -o ./dist/main.css --watch

build:
	@echo "Building for prod"
	@templ generate
	@npx tailwind -i ./style/main.css -o ./dist/main.css
	APP_ENV="pro" go run cmd/main.go
