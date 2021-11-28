.PHONY: protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/pb/user/user.proto

.PHONY: build
build:
	make protoc
	docker-compose build

.PHONY: clean
clean:
	rm -f api/pb/user/user.pb.go
	rm -f api/pb/user/user_grpc.pb.go

.PHONY: help
help:
	@echo "protoc: run protoc commands"
	@echo "build: build this project"
	@echo "clean: delete all files created by make"
