package telegram

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ravil23/baristaschool/telegrambot/entity"
	"github.com/ravil23/baristaschool/telegrambot/postgres"
)

const (
	retryPeriod     = time.Second
	maxRetriesCount = 30
)

var helpText = strings.Join([]string{
	"<b>Information</b>",
	"/progress - Show current progress",
	"/help - Show this description of usage",
	"",
	"<b>Tap to /next for getting new poll</b>",
}, "\n")

type Bot struct {
	api API
}

func NewBot() *Bot {
	return &Bot{}
}

func (b *Bot) Init() {
	log.Printf("Bot is initializing...")
	conn := postgres.NewConnection()
	botToken := GetBotTokenOrPanic()
	for i := 1; i <= maxRetriesCount; i++ {
		if api, err := NewAPI(botToken, conn); err != nil {
			log.Printf("Attempt %d failed: %v", i, err)
			time.Sleep(retryPeriod)
		} else {
			b.api = api
			log.Printf("Bot successfully initialized")
			return
		}
	}
	log.Panic("max retries count exceeded")
}

func (b *Bot) HealthCheck() {
	go func() {
		address := ":8080"
		path := "/healthcheck"
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s request to %s%s with User-Agent: %s", r.Method, r.Host, r.URL, r.UserAgent())
			_, _ = fmt.Fprint(w, `{"status": "ok"}`)
		})
		log.Printf("Listening health check on address %s%s", address, path)
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Panic(err)
		}
	}()
}

func (b *Bot) Run() {
	defer func() {
		if r := recover(); r != nil {
			b.api.SendAlert(fmt.Sprintf("Recovered from panic: %s", r))
		}
	}()
	b.api.SetMessagesHandler(func(message *entity.Message) error {
		switch message.Text {
		case "/help", "/start":
			b.api.SendHTMLMessage(message.ChatID, helpText)
			return nil
		case "/progress":
			b.api.SendProgress(message.User)
			return nil
		}
		return b.api.SendNextPoll(message.User)
	})
	b.api.SetPollAnswersHandler(func(user *entity.User, _ *entity.PollAnswer) error {
		return b.api.SendNextPoll(user)
	})
	b.serve()
}

func (b *Bot) serve() {
	go func() {
		b.api.SendAlert(fmt.Sprintf("%s started", botMention))
		if err := b.api.ListenUpdates(); err != nil {
			log.Panic(err)
		}
	}()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	b.api.SendAlert(fmt.Sprintf("%s stopped", botMention))
}
