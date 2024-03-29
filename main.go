package main

import (
	"github.com/jordanleven/reaction-bot/reactionbot"
)

var registeredReactions = reactionbot.RegisteredReactions{
	"gem": {
		Name:    "Gem",
		Channel: "gems",
	},
	"til": {
		Name:    "Today I learned",
		Channel: "til",
	},
	"test_tube": {
		Name:    "Reaction Bot Testing",
		Channel: "reaction-bot-testing",
	},
}

func main() {
	slackTokenApp := reactionbot.GetSlackTokenApp("SLACK_TOKEN_APP")
	slackTokenBot := reactionbot.GetSlackTokenBot("SLACK_TOKEN_BOT")
	registrationOptions := reactionbot.RegistrationOptions{
		SlackTokenApp:   slackTokenApp,
		SlackTokenBot:   slackTokenBot,
		RegisteredEmoji: registeredReactions,
	}
	reactionbot.New(registrationOptions)
}
