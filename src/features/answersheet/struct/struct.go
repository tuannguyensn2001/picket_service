package answersheet_struct

import "picket/src/entities"

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
