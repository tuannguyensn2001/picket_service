package test_struct

type MultipleChoiceAnswer struct {
	TestMultipleChoiceId int
	Answer               string
	Score                float64
	Type                 int32
}

type TestMultipleChoice struct {
	FilePath string
	Score    float64
	Answers  []MultipleChoiceAnswer
}

type CreateTestContentInput struct {
	TestId         int
	Typeable       int
	MultipleChoice *TestMultipleChoice
}
