FROM golang:1.17.3-alpine AS builder

WORKDIR /

COPY ./go.mod ./cmd/auth/
COPY ./go.sum ./cmd/auth/

WORKDIR /cmd/auth/

COPY ./cmd/auth/ .

RUN go build -o /bin/auth

FROM alpine:3.15.0

COPY --from=builder /bin/auth /bin/auth

ENTRYPOINT [ "/bin/auth" ]
