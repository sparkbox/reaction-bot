FROM golang:latest

RUN mkdir /reaction-bot
WORKDIR /reaction-bot
COPY . .
RUN go build main.go

CMD ["/reaction-bot/main"]
