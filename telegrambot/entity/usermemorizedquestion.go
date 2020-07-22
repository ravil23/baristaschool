package entity

import (
	"time"
)

type UserMemorizedQuestion struct {
	tableName struct{} `pg:"userMemorizedQuestion"`

	Timestamp         time.Time `pg:"timestamp,pk"`
	UserID            UserID    `pg:"user_id,pk"`
	Question          Question  `pg:"question,pk"`
	CorrectlyAnswered bool      `pg:"correctly_answered"`
}

func NewUserMemorizedQuestion(userID UserID, question Question, correctlyAnswered bool) *UserMemorizedQuestion {
	return &UserMemorizedQuestion{
		Timestamp:         time.Now(),
		UserID:            userID,
		Question:          question,
		CorrectlyAnswered: correctlyAnswered,
	}
}
