[![Build and Publish](https://github.com/ariel17/twitter-echo-bot/actions/workflows/main.yml/badge.svg)](https://github.com/ariel17/twitter-echo-bot/actions/workflows/main.yml)

# Twitter echo bot

This is a very simple application, a proof of concept as working code written in
Golang, that access the Twitter API to send an answer in the name of the account
owner for those tweets that matches the query pattern.

## Build image

```bash
$ docker build -t twitter-echo-bot .
```

## Run the application

```bash
$ docker run -p 8080:8080 \
    -e CONSUMER_API_KEY="key" \
    -e CONSUMER_API_SECRET="secret" \
    -e ACCESS_TOKEN="token" \
    -e ACCESS_TOKEN_SECRET="token_secret" \
    -e SEARCH_QUERY="#ariel17echobot" \
    -e RESPONSE_TEXT="bip bop bip!" \
    -e JOB_SECONDS=10 \
    twitter-echo-bot
```
