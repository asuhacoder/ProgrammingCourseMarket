FROM golang:1.17.3-alpine AS builder

WORKDIR /

COPY . .

COPY ./cmd/auth/ .

RUN go build -o /bin/auth

FROM alpine:3.15.0

COPY --from=builder /bin/auth /bin/auth

ENTRYPOINT [ "/bin/auth" ]
