package models

import "github.com/google/uuid"

type DeutschAdjektiv struct {
	Oid             uuid.UUID `json:"oid"`
	Type            string    `json:"type"`
	Gender          string    `json:"gender"`
	Case            string    `json:"case"`
	ArticleEnding   string    `json:"article_ending"`
	AdjectiveEnding string    `json:"adjective_ending"`
}
