init:
	cd frontend && npm ci

build-front:
	cd frontend && npm run build
	mkdir -p backend/resources
	rm -rf backend/resources/*
	cp -r frontend/build/* -t backend/resources