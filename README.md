# reverse-proxy-golang

## set up .env file with following configs

    HOST="YOUR LOCAL HOST TO PROXY AT"

## build for raspberry pi ARM v6 & v7

    env GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=0 go build -ldflags="-s -w" -o main-ARMv6 . && env GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -ldflags="-s -w" -o main-ARMv7 .
