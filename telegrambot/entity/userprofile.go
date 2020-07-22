package entity

const (
	fineCoefficientForCorrect = 0.1
	fineCoefficientForMistake = 10
)

type UserProfile struct {
	userID                     UserID
	correctlyAnsweredQuestion  map[Question]int
	mistakenlyAnsweredQuestion map[Question]int
}

func NewUserProfile(userID UserID) *UserProfile {
	return &UserProfile{
		userID:                     userID,
		correctlyAnsweredQuestion:  make(map[Question]int),
		mistakenlyAnsweredQuestion: make(map[Question]int),
	}
}

func (p *UserProfile) AddCorrectlyAnsweredQuestion(question Question) {
	p.correctlyAnsweredQuestion[question]++
}

func (p *UserProfile) AddMistakenlyAnsweredQuestion(question Question) {
	p.mistakenlyAnsweredQuestion[question]++
}

func (p *UserProfile) GetMemorizationWeight(question Question) float64 {
	correctAnswers := p.correctlyAnsweredQuestion[question]
	mistakeAnswers := p.mistakenlyAnsweredQuestion[question]
	diff := float64(correctAnswers - mistakeAnswers)
	if diff == 0 {
		return 1
	} else if diff < 0 {
		return fineCoefficientForMistake * -diff
	} else {
		return fineCoefficientForCorrect * 1 / (1 + diff)
	}
}

func (p *UserProfile) IsCorrectMemorized(question Question) bool {
	correctAnswers := p.correctlyAnsweredQuestion[question]
	mistakeAnswers := p.mistakenlyAnsweredQuestion[question]
	return correctAnswers > mistakeAnswers
}
