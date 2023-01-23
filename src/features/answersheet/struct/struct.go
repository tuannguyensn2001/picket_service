package answersheet_struct

import (
	"picket/src/entities"
	"time"
)

type StartOutput struct {
	Test    *entities.Test
	Content *entities.TestContent
}

type UserAnswerInput struct {
	TestId         int    `validate:"required"`
	QuestionId     int    `validate:"required"`
	Answer         string `validate:"required"`
	PreviousAnswer string
}

type GetContentOutput struct {
	Content  *entities.TestContent
	TimeLeft *time.Duration
}
