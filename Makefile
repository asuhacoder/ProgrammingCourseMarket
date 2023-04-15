.PHONY: protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/pkg/pb/auth/auth.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/pkg/pb/user/user.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/pkg/pb/course/course.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/pkg/pb/lesson/lesson.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/pkg/pb/case/case.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/pkg/pb/runner/runner.proto

.PHONY: build
build:
	make protoc
	docker compose build

.PHONY: deploy
deploy:
	./scripts/push-image.sh

.PHONY: clean
clean:
	rm -f internal/pkg/pb/auth/auth.pb.go
	rm -f internal/pkg/pb/auth/auth_grpc.pb.go
	rm -f internal/pkg/pb/user/user.pb.go
	rm -f internal/pkg/pb/user/user_grpc.pb.go
	rm -f internal/pkg/pb/course/course.pb.go
	rm -f internal/pkg/pb/course/course_grpc.pb.go
	rm -f internal/pkg/pb/lesson/lesson.pb.go
	rm -f internal/pkg/pb/lesson/lesson_grpc.pb.go
	rm -f internal/pkg/pb/runner/runner.pb.go
	rm -f internal/pkg/pb/runner/runner_grpc.pb.go

.PHONY: help
help:
	@echo "protoc: run protoc commands"
	@echo "build: build this project"
	@echo "deploy: deploy this project"
	@echo "clean: delete all files created by make"
