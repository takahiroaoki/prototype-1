init:
	cd frontend && npm ci
	cd backend && go mod tidy

build-front:
	cd frontend && npm run build
	mkdir -p backend/resources
	rm -rf backend/resources/*
	cp -r frontend/build/* -t backend/resources

run-back:
	cd backend && go run main.go

run: build-front run-back