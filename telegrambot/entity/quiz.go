package entity

import (
	"math"
	"math/rand"
)

type Question string

func (q Question) String() string {
	return string(q)
}

type Answer struct {
	CorrectOption string
	InvalidOptions []string
}

type Quiz struct {
	tests map[Question]Answer
}

func NewEmptyQuiz() *Quiz {
	return &Quiz{
		tests: make(map[Question]Answer),
	}
}

func NewQuiz(tests map[Question]Answer) *Quiz {
	return &Quiz{
		tests: tests,
	}
}

func (q *Quiz) GetRandomQuestion() Question {
	allQuestions := make([]Question, 0, len(q.tests))
	for question := range q.tests{
		allQuestions = append(allQuestions, question)
	}
	return allQuestions[rand.Intn(len(allQuestions))]
}

func (q *Quiz) GetQuestionByUserProfile(userProfile *UserProfile) (Question, float64, bool) {
	weights := make(map[Question]float64, len(q.tests))
	weightsSum := 0.
	weightsMax := 0.
	for question := range q.tests {
		weight := userProfile.GetMemorizationWeight(question)
		weights[question] = weight
		weightsSum += weight
		weightsMax = math.Max(weightsMax, weight)
	}
	randomPoint := rand.Float64() * weightsSum
	allQuestionsMemorized := weightsMax < 1
	var left, right float64
	for question, weight := range weights {
		right = left + weight
		if left <= randomPoint && randomPoint < right {
			return question, weight, allQuestionsMemorized
		}
		left = right
	}
	return "", 0, false
}

func (q *Quiz) GetCorrectMemorizedQuestionsCount(userProfile *UserProfile) int {
	correctMemorizedQuestionsCount := 0
	for question := range q.tests {
		if userProfile.IsCorrectMemorized(question) {
			correctMemorizedQuestionsCount++
		}
	}
	return correctMemorizedQuestionsCount
}

func (q *Quiz) GetAnswer(question Question) Answer {
	return q.tests[question]
}

func (q *Quiz) GetQuestionsCount() int {
	return len(q.tests)
}

func (q *Quiz) Update(other *Quiz) *Quiz {
	for question, translations := range other.tests {
		q.tests[question] = translations
	}
	return q
}
