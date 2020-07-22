package telegram

import (
	"log"
	"time"

	"github.com/ravil23/baristaschool/telegrambot/dao"
	"github.com/ravil23/baristaschool/telegrambot/entity"
	"github.com/ravil23/baristaschool/telegrambot/postgres"
)

const (
	userMemorizedQuestionsTTL = -2 * 7 * 24 * time.Hour
)

var alreadyFinishedUsers = map[entity.UserID]bool{}

type UserProfileManager struct {
	pollsStates  map[entity.PollID]*entity.Poll
	userProfiles map[entity.UserID]*entity.UserProfile

	userDAO                  dao.UserDAO
	userMemorizedQuestionDAO dao.UserMemorizedQuestionDAO
}

func NewUserProfileManager(conn *postgres.Connection, userDAO dao.UserDAO) (*UserProfileManager, error) {
	userMemorizedQuestionDAO, err := dao.NewUserMemorizedQuestionDAO(conn)
	if err != nil {
		return nil, err
	}
	m := &UserProfileManager{
		pollsStates:              make(map[entity.PollID]*entity.Poll),
		userProfiles:             make(map[entity.UserID]*entity.UserProfile),
		userDAO:                  userDAO,
		userMemorizedQuestionDAO: userMemorizedQuestionDAO,
	}
	go m.initUserProfiles()
	return m, nil
}

func (m *UserProfileManager) AddPoll(poll *entity.Poll) {
	m.pollsStates[poll.ID] = poll
}

func (m *UserProfileManager) AddPollAnswer(userID entity.UserID, pollAnswer *entity.PollAnswer) error {
	poll, found := m.pollsStates[pollAnswer.PollID]
	if !found {
		log.Printf("Poll corresponded to answer is not found: %+v", pollAnswer)
		return nil
	}
	defer delete(m.pollsStates, pollAnswer.PollID)
	correctlyAnswered := poll.AllIsCorrect(pollAnswer.ChosenOptions)
	userMemorizedQuestion := entity.NewUserMemorizedQuestion(userID, poll.Question, correctlyAnswered)
	if err := m.userMemorizedQuestionDAO.Upsert(userMemorizedQuestion); err != nil {
		return err
	}

	m.updateUserProfiles(userID, poll.Question, correctlyAnswered)
	return nil
}

func (m *UserProfileManager) GetUserProfile(userID entity.UserID) (*entity.UserProfile, bool) {
	userProfile, found := m.userProfiles[userID]
	return userProfile, found
}

func (m *UserProfileManager) initUserProfiles() {
	users, err := m.userDAO.FindAll()
	if err != nil {
		panic(err)
	}
	log.Printf("Found users count: %d", len(users))
	from := time.Now().Add(userMemorizedQuestionsTTL)
	for _, user := range users {
		log.Printf("Init profile for user: %+v", user)
		userMemorizedQuestions, err := m.userMemorizedQuestionDAO.FindByUserID(user.ID, from)
		if err != nil {
			panic(err)
		}
		log.Printf("User %d has memorized %d questions for last %s", user.ID, len(userMemorizedQuestions), userMemorizedQuestionsTTL)
		for _, userMemorizedQuestion := range userMemorizedQuestions {
			m.updateUserProfiles(user.ID, userMemorizedQuestion.Question, userMemorizedQuestion.CorrectlyAnswered)
		}
	}
}

func (m *UserProfileManager) updateUserProfiles(userID entity.UserID, question entity.Question, correctlyAnswered bool) {
	if _, found := m.userProfiles[userID]; !found {
		m.userProfiles[userID] = entity.NewUserProfile(userID)
	}
	userProfile := m.userProfiles[userID]
	if correctlyAnswered {
		userProfile.AddCorrectlyAnsweredQuestion(question)
	} else {
		userProfile.AddMistakenlyAnsweredQuestion(question)
	}
}
