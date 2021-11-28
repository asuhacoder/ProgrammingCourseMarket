FROM golang:1.17.3-alpine AS builder

WORKDIR /go/src/github.com/Asuha-a/ProgrammingCourseMarket/api/

COPY ./go.mod ./user/
COPY ./go.sum ./user/
RUN mkdir -p ./user/pb/user/
COPY ./pb/user/. ./user/pb/user/

WORKDIR /go/src/github.com/Asuha-a/ProgrammingCourseMarket/api/user/
COPY ./user/. .

RUN go build -o /bin/user

FROM alpine:3.15.0

COPY --from=builder /bin/user /bin/user

ENTRYPOINT [ "/bin/user" ]