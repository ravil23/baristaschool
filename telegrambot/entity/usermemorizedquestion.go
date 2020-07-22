package entity

import (
	"time"
)

type UserMemorizedQuestion struct {
	tableName struct{} `pg:"userMemorizedQuestion"`

	Timestamp           time.Time `pg:"timestamp,pk"`
	UserID              UserID    `pg:"user_id,pk"`
	Question            Question      `pg:"question,pk"`
	CorrectlyTranslated bool      `pg:"correctly_translated"`
}

func NewUserMemorizedQuestion(userID UserID, question Question, correctlyTranslated bool) *UserMemorizedQuestion {
	return &UserMemorizedQuestion{
		Timestamp:           time.Now(),
		UserID:              userID,
		Question:            question,
		CorrectlyTranslated: correctlyTranslated,
	}
}
