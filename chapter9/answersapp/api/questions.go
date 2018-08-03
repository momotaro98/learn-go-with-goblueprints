package api

import (
	"errors"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Question struct {
	Key          *datastore.Key `json:"id" datastore:"-"`
	CTime        time.Time      `json:"created"`
	Question     string         `json:"question"`
	User         UserCard       `json:"user"`
	AnswersCount int            `json:"answers_count"`
}

// OK checks if a Question have no problem.
func (q Question) OK() error {
	if len(q.Question) < 10 {
		return errors.New("question is too short")
	}
	return nil
}

// Create creates a new Question into Datastore
func (q *Question) Create(ctx context.Context) error {
	log.Debugf(ctx, "Saving question: %s", q.Question)
	if q.Key == nil {
		q.Key = datastore.NewIncompleteKey(ctx, "Question", nil)
	}
	user, err := UserFromAEUser(ctx)
	if err != nil {
		return nil
	}
	q.User = user.Card()
	q.CTime = time.Now()
	q.Key, err = datastore.Put(ctx, q.Key, q)
	if err != nil {
		return err
	}
	return nil
}

// Update updates a Question on Datastore.
func (q *Question) Update(ctx context.Context) error {
	if q.Key == nil {
		q.Key = datastore.NewIncompleteKey(ctx, "Question", nil)
	}
	var err error
	q.Key, err = datastore.Put(ctx, q.Key, q)
	if err != nil {
		return err
	}
	return nil
}

// GetQuestion retrieves a Question from Datastore with a key.
func GetQuestion(ctx context.Context, key *datastore.Key) (*Question, error) {
	var q Question
	err := datastore.Get(ctx, key, &q)
	if err != nil {
		return nil, err
	}
	q.Key = key
	return &q, nil
}

// TopQuestions returns questions from top order.
func TopQuestions(ctx context.Context) ([]*Question, error) {
	var questions []*Question
	questionKeys, err := datastore.NewQuery("Question").
		Order("-AnswersCount").
		Order("-CTime").
		Limit(25).
		GetAll(ctx, questions)
	if err != nil {
		return nil, err
	}
	for i := range questions {
		questions[i].Key = questionKeys[i]
	}
	log.Debugf(ctx, "questions: %s", questions)
	return questions, nil
}
