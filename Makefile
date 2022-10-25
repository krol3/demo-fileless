.PHONY: build
build :
	@echo "Building ...."
	CGO_ENABLED=0 go build $(LDFLAGS) -o ./bin/demo-memfd ./pkg/server.go

build-image:
	docker build -t krol/demo-memfd:v1.1 .