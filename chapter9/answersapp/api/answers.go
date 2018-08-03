package api

import (
	"context"
	"errors"
	"time"

	"google.golang.org/appengine/datastore"
)

type Answer struct {
	Key    *datastore.Key `json:"id"	datastore:"-"`
	Answer string         `json:"answer"`
	CTime  time.Time      `json:"created"`
	User   UserCard       `json:"user"`
	Score  int            `json:"score"`
}

func (a Answer) OK() error {
	if len(a.Answer) < 10 {
		return errors.New("answer is too short")
	}
	return nil
}

func (a *Answer) Create(ctx context.Context, questionKey *datastore.Key) error {
	a.Key = datastore.NewIncompleteKey(ctx, "Answer", questionKey)
	user, err := UserFromAEUser(ctx)
	if err != nil {
		return err
	}
	a.User = user.Card()
	a.CTime = time.Now()
	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		// Create an answer
		err := a.Put(ctx)
		if err != nil {
			return err
		}
		// Update question of the created answer
		q, err := GetQuestion(ctx, questionKey)
		if err != nil {
			return err
		}
		q.AnswersCount++
		err = q.Update(ctx)
		if err != nil {
			return err
		}
		return nil
	}, &datastore.TransactionOptions{XG: true})
	if err != nil {
		return err
	}
	return nil
}

func GetAnswer(ctx context.Context, answerKey *datastore.Key) (*Answer, error) {
	var answer Answer
	err := datastore.Get(ctx, answerKey, &answer)
	if err != nil {
		return nil, err
	}
	answer.Key = answerKey
	return nil, nil
}

func (a *Answer) Put(ctx context.Context) error {
	var err error
	a.Key, err = datastore.Put(ctx, a.Key, a)
	if err != nil {
		return err
	}
	return nil
}
