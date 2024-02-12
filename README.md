# Reaction Bot

A small but mighty implementation of the [Reaction Bot library][reaction_bot_repo] to post messages to channels based on emoji reactions.

## Running locally

To run this project locally, copy the contents of `.env-sample` to `.env` and replace the token with the correct tokens from Slack. You can either run the project with a local Go installation, or you can use Docker.

### Go

To install dependencies, run:

```cli
go get
```

To initialize, run:

```cli
go run main.go
```

To compile the binary, run:

```cli
go build main.go
```

### Docker

To build the container after initial setup or making changes, run:

```cli
docker compose build
```

To run the binary from the container, run:

```cli
docker compose up
```

## Customizing

Looking to shake things up? You can customize the emoji and the channels they should post to in `reaction-bot.go` by modifying `registeredReactions`. For configuration options on registering emoji, see the [docs][reaction_bot_register_emoji].

[reaction_bot_repo]: http://github.com/jordanleven/reaction-bot
[reaction_bot_register_emoji]: https://github.com/jordanleven/reaction-bot#registering-emoji

## Deploying

When changes are pushed to `main`, an image is built and published to the GitHub Container Registry. From the server, you can pull the image and restart the bot with following commands:

```cli
docker pull ghcr.io/sparkbox/reaction-bot
docker compose up -d
```
