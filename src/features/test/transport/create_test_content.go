package test_transport

import (
	"context"
	"errors"
	"gorm.io/gorm"
	test_struct "myclass_service/src/features/test/struct"
	errpkg "myclass_service/src/packages/err"
	testpb "myclass_service/src/pb/test"
)

func (t *transport) CreateContent(ctx context.Context, request *testpb.CreateTestContentRequest) (*testpb.CreateTestContentResponse, error) {
	input := test_struct.CreateTestContentInput{
		TestId:   int(request.TestId),
		Typeable: int(request.Typeable),
		MultipleChoice: &test_struct.TestMultipleChoice{
			FilePath: request.MultipleChoice.FilePath,
			Score:    float64(request.MultipleChoice.Score),
		},
	}
	answers := make([]test_struct.MultipleChoiceAnswer, len(request.MultipleChoice.Answers))
	for index, item := range request.MultipleChoice.Answers {
		answers[index] = test_struct.MultipleChoiceAnswer{
			Answer: item.Answer,
			Score:  float64(item.Score),
			Type:   item.Type,
		}
	}
	input.MultipleChoice.Answers = answers

	err := t.usecase.CreateContent(ctx, input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errpkg.General.NotFound)
		}
		panic(err)
	}

	resp := testpb.CreateTestContentResponse{
		Message: "success",
	}
	return &resp, nil
}
