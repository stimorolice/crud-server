FROM golang:1.23-alpine AS builder

COPY . /root/bruh
WORKDIR /root/bruh

RUN go mod download
RUN go build -o /bin/crud_server cmd/server/main.go


FROM alpine:latest


WORKDIR /root/

COPY --from=builder /bin/crud_server .

CMD [ "/root/crud_server" ]