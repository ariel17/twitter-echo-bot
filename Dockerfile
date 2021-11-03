FROM golang:stretch AS test
WORKDIR /test
ENV CGO_ENABLED=1
COPY . .
RUN go test -race -v ./...

FROM golang:alpine AS build
WORKDIR /build
ENV CGO_ENABLED=0
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o app .


FROM alpine:latest
WORKDIR /app
COPY --from=build /build/app .

ENV ENVIRONMENT="production"

ENV CONSUMER_API_KEY="api_key"
ENV CONSUMER_API_SECRET="api_secret"
ENV ACCESS_TOKEN="access_token"
ENV ACCESS_TOKEN_SECRET="access_token_secret"

ENV SEARCH_QUERY="#ariel17echobot"
ENV RESPONSE_TEXT="Hola! Esta respuesta está automatizada. https://github.com/ariel17/twitter-echo-bot/"
ENV JOB_SECONDS=10

EXPOSE 8080

CMD ["./app"]
