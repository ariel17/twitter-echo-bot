[![Build and Publish](https://github.com/ariel17/twitter-echo-bot/actions/workflows/main.yml/badge.svg)](https://github.com/ariel17/twitter-echo-bot/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ariel17/twitter-echo-bot)](https://goreportcard.com/report/github.com/ariel17/twitter-echo-bot)

# Twitter echo bot

This is a very simple application, a proof of concept as working code written in
Golang, that access Twitter API to send an answer in the name of the account
owner for those tweets that matches the query pattern.

## Example

[See it on Twitter.](https://twitter.com/ariel_17_/status/1451647851180740609)

![automated response](./docs/example.png)

## Docker image

The Docker image is hosted at [Docker Hub](https://hub.docker.com/r/ariel17/twitter-echo-bot)
and it is build and pushed through [GitHub Actions](./actions).

### How to build it

```bash
$ docker build -t twitter-echo-bot .
```

### How to execute it

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

## Production deployment

The deployment is made using [Terraform](https://www.terraform.io/) to a AWS
t2.nano instance.

TODO: not yet completed.

### Development lifecycle

![Services integration diagram](./docs/lifecycle.png)
