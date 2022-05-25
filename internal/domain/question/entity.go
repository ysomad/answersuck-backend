package question

import "time"

type Question struct {
	Id             int       `json:"id"`
	Text           string    `json:"text"`
	Answer         string    `json:"answer"`
	AnswerImageURL *string   `json:"answerImage"`
	Author         string    `json:"author"`
	MediaURL       *string   `json:"media"`
	MediaType      *string   `json:"mediaType"`
	LanguageId     int       `json:"languageId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
