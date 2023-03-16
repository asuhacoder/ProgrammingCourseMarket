FROM golang:1.17.3-alpine AS builder

WORKDIR /

COPY . .

COPY ./cmd/case/ .

RUN go build -o /bin/test_case

FROM alpine:3.15.0

COPY --from=builder /bin/test_case /bin/test_case

ENTRYPOINT [ "/bin/test_case" ]
