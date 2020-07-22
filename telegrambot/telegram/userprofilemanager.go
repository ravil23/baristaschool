package telegram

import (
	"log"
	"time"

	"github.com/ravil23/baristaschool/telegrambot/dao"
	"github.com/ravil23/baristaschool/telegrambot/entity"
	"github.com/ravil23/baristaschool/telegrambot/postgres"
)

const (
	userMemorizedTermsTTL = -2 * 7 * 24 * time.Hour
)

var alreadyFinishedUsers = map[entity.UserID]bool{}

type UserProfileManager struct {
	pollsStates  map[entity.PollID]*entity.Poll
	userProfiles map[entity.UserID]*entity.UserProfile

	userDAO              dao.UserDAO
	userMemorizedTermDAO dao.UserMemorizedQuestionDAO
}

func NewUserProfileManager(conn *postgres.Connection, userDAO dao.UserDAO) (*UserProfileManager, error) {
	userMemorizedTermDAO, err := dao.NewUserMemorizedQuestionDAO(conn)
	if err != nil {
		return nil, err
	}
	m := &UserProfileManager{
		pollsStates:          make(map[entity.PollID]*entity.Poll),
		userProfiles:         make(map[entity.UserID]*entity.UserProfile),
		userDAO:              userDAO,
		userMemorizedTermDAO: userMemorizedTermDAO,
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
	correctlyTranslated := poll.AllIsCorrect(pollAnswer.ChosenOptions)
	userMemorizedTerm := entity.NewUserMemorizedQuestion(userID, poll.Question, correctlyTranslated)
	if err := m.userMemorizedTermDAO.Upsert(userMemorizedTerm); err != nil {
		return err
	}

	m.updateUserProfiles(userID, poll.Question, correctlyTranslated)
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
	from := time.Now().Add(userMemorizedTermsTTL)
	for _, user := range users {
		log.Printf("Init profile for user: %+v", user)
		userMemorizedTerms, err := m.userMemorizedTermDAO.FindByUserID(user.ID, from)
		if err != nil {
			panic(err)
		}
		log.Printf("User %d has %d memorized terms for last %s", user.ID, len(userMemorizedTerms), userMemorizedTermsTTL)
		for _, userMemorizedTerm := range userMemorizedTerms {
			m.updateUserProfiles(user.ID, userMemorizedTerm.Question, userMemorizedTerm.CorrectlyTranslated)
		}
	}
}

func (m *UserProfileManager) updateUserProfiles(userID entity.UserID, question entity.Question, correctlyTranslated bool) {
	if _, found := m.userProfiles[userID]; !found {
		m.userProfiles[userID] = entity.NewUserProfile(userID)
	}
	userProfile := m.userProfiles[userID]
	if correctlyTranslated {
		userProfile.AddCorrectlyAnsweredQuestion(question)
	} else {
		userProfile.AddMistakenlyAnsweredQuestion(question)
	}
}
