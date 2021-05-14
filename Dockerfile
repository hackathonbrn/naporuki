FROM golang:1.16-alpine AS builder

ADD api/ .
RUN unset GOPATH && go build -o /server

FROM alpine:3

CMD ["/web/server"]

WORKDIR /web
COPY --from=builder /server .