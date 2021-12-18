FROM golang:1.17.3-alpine AS builder

WORKDIR /

COPY . .

COPY ./cmd/lesson/ .

RUN go build -o /bin/lesson

FROM alpine:3.15.0

COPY --from=builder /bin/lesson /bin/lesson

ENTRYPOINT [ "/bin/lesson" ]
