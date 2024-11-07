init:
	cd frontend && sudo chmod -R 777 node_modules && npm ci
	cd backend && go mod tidy

build-front:
	cd frontend && npm run build
	mkdir -p backend/resources
	rm -rf backend/resources/*
	cp -r frontend/build/* -t backend/resources

build-back:
	cd backend && CGO_ENABMED=0 GOOS=linux GOARCH=amd64 go build -o prototype-1

run-back:
	cd backend && go run main.go

run: build-front run-back

build: build-front build-back