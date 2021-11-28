FROM golang:1.17.3-alpine AS builder

WORKDIR /go/src/github.com/Asuha-a/ProgrammingCourseMarket/api

COPY . .

RUN go build -o /bin/gateway

FROM alpine:3.15.0

COPY --from=builder /bin/gateway /bin/gateway

ENTRYPOINT [ "/bin/gateway" ]
