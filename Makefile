BUILD_SUCCESS := The build is complete!

test:
	@go test -race -cover ./... | grep -v '^?'
fmt:
	go fmt ./...
init-env:
	echo "ADMIN_SECRET_KEY=$(tr -dc 'a-z0-9' </dev/urandom | head -c 32)" > .env && echo "JWT_SECRET_KEY=$(tr -dc 'a-z0-9' </dev/urandom | head -c 32)" >> .env
build:
	npm install --prefix web && npm run build --prefix web && go build ./cmd/main.go && echo $(BUILD_SUCCESS)
count-lines:
	@echo "total code lines:" && find . -name "*.go" -exec cat {} \; | wc -l
