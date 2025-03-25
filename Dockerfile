# STEP-1
# build app from source

FROM golang:1.24.1-alpine3.21 AS builder

WORKDIR /mysource

COPY ./cmd ./cmd
# COPY ./vendor* ./vendor
COPY ./go.mod ./go.sum ./.env ./
COPY ./internal ./internal  
COPY ./database ./database
COPY ./input.txt ./

RUN go build -o app ./cmd/main.go

# STEP-2
# make container

FROM alpine:3.21

WORKDIR /myapp

COPY --from=builder /mysource ./

CMD [ "/myapp/app" ]
