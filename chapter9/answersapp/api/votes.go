package api

import (
	"time"

	"google.golang.org/appengine/datastore"
)

type Vote struct {
	Key      *datastore.Key `json:"id" datastore:"-"`
	MTime    time.Time      `json:"last_modified" datastore:",noindex"`
	Question QuestionCard   `json:"question" datastore:",noindex"`
	Answer   AnswerCard     `json:"answer" datastore:",noindex"`
	User     UserCard       `json:"user" datastore:",noindex"`
	Score    int            `json:"score" datastore:",noindex"`
}
