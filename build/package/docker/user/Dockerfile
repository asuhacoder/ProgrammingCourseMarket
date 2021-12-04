FROM golang:1.17.3-alpine AS builder

WORKDIR /

COPY . .

COPY ./cmd/user/ .

RUN go build -o /bin/user

FROM alpine:3.15.0

COPY --from=builder /bin/user /bin/user

ENTRYPOINT [ "/bin/user" ]
