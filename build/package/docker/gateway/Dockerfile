FROM golang:1.17.3-alpine AS builder

WORKDIR /

COPY . .

COPY ./cmd/gateway/ .

RUN go build -o /bin/gateway

FROM alpine:3.15.0

COPY --from=builder /bin/gateway /bin/gateway

ENTRYPOINT [ "/bin/gateway" ]
