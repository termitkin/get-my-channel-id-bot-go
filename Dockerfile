FROM golang:1.20-alpine as build

WORKDIR /app

COPY app ./app
COPY go.mod ./

WORKDIR /app/app

RUN go build -o /get-my-channel-id

FROM alpine:latest 

COPY --from=build /get-my-channel-id ./

CMD [ "/get-my-channel-id" ]
