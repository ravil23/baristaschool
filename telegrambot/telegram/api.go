package telegram

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/baristaschool/telegrambot/collection"
	"github.com/ravil23/baristaschool/telegrambot/dao"
	"github.com/ravil23/baristaschool/telegrambot/entity"
	"github.com/ravil23/baristaschool/telegrambot/postgres"
)

const (
	timeout             = 10
	alertsChatID        = -1001215763168
	botNickName         = "BaristaSchoolBot"
	botMention          = "@" + botNickName
)

type API interface {
	SetMessagesHandler(handlerFunc func(*entity.Message) error)
	SetPollAnswersHandler(handlerFunc func(*entity.User, *entity.PollAnswer) error)
	ListenUpdates() error
	SendNextPoll(user *entity.User) error
	SendAlert(text string)
	SendMessage(chatID entity.ChatID, text string)
	SendHTMLMessage(chatID entity.ChatID, text string)
	SendProgress(user *entity.User)
}

var _ API = (*api)(nil)

type api struct {
	hostName string

	tgAPI     *tgbotapi.BotAPI
	tgUpdates tgbotapi.UpdatesChannel

	userDAO dao.UserDAO

	userProfileManager *UserProfileManager

	messagesHandler    func(update *tgbotapi.Update) error
	pollAnswersHandler func(update *tgbotapi.Update) error
}

func NewAPI(botToken string, conn *postgres.Connection) (*api, error) {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "unknown_host"
	}

	tgAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout
	tgUpdates := tgAPI.GetUpdatesChan(u)

	userDAO, err := dao.NewUserDAO(conn)
	if err != nil {
		return nil, err
	}

	userProfileManager, err := NewUserProfileManager(conn, userDAO)
	if err != nil {
		return nil, err
	}

	return &api{
		hostName:  hostName,
		tgUpdates: tgUpdates,
		tgAPI:     tgAPI,

		userDAO: userDAO,

		userProfileManager: userProfileManager,
	}, nil
}

func (api *api) SetMessagesHandler(handlerFunc func(*entity.Message) error) {
	api.messagesHandler = func(update *tgbotapi.Update) error {
		user := entity.NewUser(update.Message.From)
		user.ChatID = entity.ChatID(update.Message.Chat.ID)
		if err := api.userDAO.Upsert(user); err != nil {
			return err
		}

		if update.Message.Command() == "start" {
			api.SendAlert(fmt.Sprintf("%s started conversation with %s", user.GetFormattedName(), botMention))
		}

		message := entity.NewMessage(update.Message, user)

		if err := handlerFunc(message); err != nil {
			return err
		}
		return nil
	}
}

func (api *api) SetPollAnswersHandler(handlerFunc func(*entity.User, *entity.PollAnswer) error) {
	api.pollAnswersHandler = func(update *tgbotapi.Update) error {
		user, err := api.userDAO.Find(entity.UserID(update.PollAnswer.User.ID))
		if err != nil {
			return err
		}

		pollAnswer := entity.NewPollAnswer(update.PollAnswer)

		if err := api.userProfileManager.AddPollAnswer(user.ID, pollAnswer); err != nil {
			return err
		}

		return handlerFunc(user, pollAnswer)
	}
}

func (api *api) ListenUpdates() error {
	for update := range api.tgUpdates {
		log.Printf("Handle update: %+v", update)
		if update.Message != nil {
			if err := api.messagesHandler(&update); err != nil {
				return err
			}
		}
		if update.PollAnswer != nil {
			if err := api.pollAnswersHandler(&update); err != nil {
				return err
			}
		}
	}
	return nil
}

func (api *api) SendNextPoll(user *entity.User) error {
	poll, found := api.getNextPoll(user)
	if !found {
		headerText := fmt.Sprintf("<b>Congratulations!</b>\nYou have answered correctly to all questions.")
		progressText := api.getProgressByUser(user)
		text := strings.Join([]string{
			headerText,
			progressText,
			"",
			"Type /next to continue learning.",
		}, "\n")
		api.SendHTMLMessage(user.ChatID, text)
		return nil
	}
	tgPoll := poll.ToChatable(user.ChatID)
	tgMessage, err := api.tgAPI.Send(tgPoll)
	if err != nil {
		return err
	}
	if tgMessage.Poll == nil {
		return fmt.Errorf("returned message does not contain poll: %+v", tgMessage)
	}
	poll.ID = entity.PollID(tgMessage.Poll.ID)
	api.userProfileManager.AddPoll(poll)
	return nil
}

func (api *api) getNextPoll(user *entity.User) (*entity.Poll, bool) {
	quiz := collection.Quiz
	var question entity.Question
	var weight float64
	var finished bool
	if userProfile, found := api.userProfileManager.GetUserProfile(user.ID); found {
		question, weight, finished = quiz.GetQuestionByUserProfile(userProfile)
		if finished && !alreadyFinishedUsers[user.ID] {
			alreadyFinishedUsers[user.ID] = true
			return nil, false
		}
	} else {
		question = quiz.GetRandomQuestion()
	}
	answer := quiz.GetAnswer(question)
	poll := &entity.Poll{
		Question: question,
		Weight:   weight,
		Type:     entity.PollTypeQuiz,
		IsPublic: true,
		Options: []*entity.PollOption{
			{Value: answer.CorrectOption, IsCorrect: true},
		},
	}
	for _, option := range answer.InvalidOptions {
		poll.Options = append(poll.Options, &entity.PollOption{
			Value: option,
			IsCorrect:   false,
		})
	}
	rand.Shuffle(len(poll.Options), func(i, j int) {
		poll.Options[i], poll.Options[j] = poll.Options[j], poll.Options[i]
	})
	return poll, true
}

func (api *api) SendProgress(user *entity.User) {
	api.SendMessage(user.ChatID, api.getProgressByUser(user))
}

func (api *api) getProgressByUser(user *entity.User) string {
	userProfile, _ := api.userProfileManager.GetUserProfile(user.ID)
	totalQuestionsCount := collection.Quiz.GetQuestionsCount()
	correctMemorizedQuestionsCount := collection.Quiz.GetCorrectMemorizedQuestionsCount(userProfile)
	return fmt.Sprintf(
		"Progress is %s (%d questions from %d memorized)",
		fmt.Sprintf("%.1f%%", 100*float64(correctMemorizedQuestionsCount)/float64(totalQuestionsCount)),
		correctMemorizedQuestionsCount,
		totalQuestionsCount,
	)
}

func (api *api) SendAlert(text string) {
	api.SendMessage(alertsChatID, fmt.Sprintf("[%s] %s", api.hostName, text))
}

func (api *api) SendMessage(chatID entity.ChatID, text string) {
	api.sendMessage(chatID, text, "")
}

func (api *api) SendHTMLMessage(chatID entity.ChatID, text string) {
	api.sendMessage(chatID, text, tgbotapi.ModeHTML)
}

func (api *api) sendMessage(chatID entity.ChatID, text string, parseMode string) {
	log.Printf("Chat ID: %d, Parse mode: %s, Text: %s", chatID, parseMode, text)
	tgMessage := tgbotapi.NewMessage(int64(chatID), text)
	tgMessage.ParseMode = parseMode
	_, err := api.tgAPI.Send(tgMessage)
	if err != nil {
		log.Printf("Error on sending message: %s", err)
	}
}

func GetBotTokenOrPanic() string {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Panic("bot token is empty")
	}
	return botToken
}
