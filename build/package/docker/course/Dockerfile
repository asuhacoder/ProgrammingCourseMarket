FROM golang:1.17.3-alpine AS builder

WORKDIR /

COPY . .

COPY ./cmd/course/ .

RUN go build -o /bin/course

FROM alpine:3.15.0

COPY --from=builder /bin/course /bin/course

ENTRYPOINT [ "/bin/course" ]
