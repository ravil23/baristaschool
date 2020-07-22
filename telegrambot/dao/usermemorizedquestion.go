package dao

import (
	"time"

	"github.com/go-pg/pg/v9/orm"

	"github.com/ravil23/baristaschool/telegrambot/entity"
	"github.com/ravil23/baristaschool/telegrambot/postgres"
)

type UserMemorizedQuestionDAO interface {
	FindByUserID(userID entity.UserID, from time.Time) ([]*entity.UserMemorizedQuestion, error)
	Upsert(userMemorizedQuestion *entity.UserMemorizedQuestion) error
}

var _ UserMemorizedQuestionDAO = (*userMemorizedQuestionDAO)(nil)

type userMemorizedQuestionDAO struct {
	conn *postgres.Connection
}

func NewUserMemorizedQuestionDAO(conn *postgres.Connection) (*userMemorizedQuestionDAO, error) {
	dao := &userMemorizedQuestionDAO{
		conn: conn,
	}
	if err := dao.ensureSchema(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (dao *userMemorizedQuestionDAO) ensureSchema() error {
	options := &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}
	return dao.conn.CreateTable((*entity.UserMemorizedQuestion)(nil), options)
}

func (dao *userMemorizedQuestionDAO) FindByUserID(userID entity.UserID, from time.Time) ([]*entity.UserMemorizedQuestion, error) {
	var userMemorizedQuestions []*entity.UserMemorizedQuestion
	err := dao.conn.Model(&userMemorizedQuestions).
		Where("user_id = ?", userID).
		Where("timestamp >= ?", from).
		Select()
	if err != nil {
		return nil, err
	}
	return userMemorizedQuestions, nil
}

func (dao *userMemorizedQuestionDAO) Upsert(userMemorizedQuestion *entity.UserMemorizedQuestion) error {
	_, err := dao.conn.Model(userMemorizedQuestion).
		OnConflict("(timestamp, user_id, question) DO NOTHING").
		Insert()
	return err
}
