FROM golang:1.17.3-alpine AS builder

WORKDIR /

COPY . .

COPY ./cmd/case/ .

RUN go build -o /bin/case

FROM alpine:3.15.0

COPY --from=builder /bin/case /bin/case

ENTRYPOINT [ "/bin/case" ]
