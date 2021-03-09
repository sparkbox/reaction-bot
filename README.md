# Reaction Bot

A small but mighty implementation of the [Reaction Bot library][reaction_bot_repo] to post messages to channels based on emoji reactions.

## Running locally

To run this project locally, copy the contents of `.env-sample` to `.env` and replace the token with the correct tokens from Slack.

To install dependencies, run:

```cli
go get
```

To initialize, run:

```cli
go run reaction-bot.go
```

To compile the binary, run:

```cli
go build reaction-bot.go
```

### Customizing

Looking to shake things up? You can customize the emoji and the channels they should post to in `reaction-bot.go` by modifying `registeredReactions`. For configuration options on registering emoji, see the [docs][reaction_bot_register_emoji].

[reaction_bot_repo]: http://github.com/jordanleven/reaction-bot
[reaction_bot_register_emoji]: https://github.com/jordanleven/reaction-bot#registering-emoji
