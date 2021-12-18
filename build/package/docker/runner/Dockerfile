FROM golang:1.17.3-alpine AS builder

WORKDIR /

COPY . .

COPY ./cmd/runner/ .

RUN go build -o /bin/runner

FROM alpine:3.15.0

COPY --from=builder /bin/runner /bin/runner

ENTRYPOINT [ "/bin/runner" ]
