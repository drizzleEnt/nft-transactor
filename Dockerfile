FROM golang:1.22.2-alpine AS builder

COPY . /github.com/drizzleent/nft-transactor/source/
WORKDIR /github.com/drizzleent/nft-transactor/source/

RUN go mod download
RUN go build -o ./bin/crud_server ./main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/drizzleent/nft-transactor/source/.env .
COPY --from=builder /github.com/drizzleent/nft-transactor/source/bin/crud_server .

EXPOSE 8080
CMD [ "./crud_server" ]