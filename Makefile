IMAGE="ariel17/twitter-echo-bot:dev"

build:
	docker build -t $(IMAGE) .

run: build
	docker run --env-file=.env -t $(IMAGE)
