package answersheet_transport

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/durationpb"
	answersheet_struct "picket/src/features/answersheet/struct"
	errpkg "picket/src/packages/err"
	answersheetpb "picket/src/pb/answer_sheet"
	"picket/src/utils"
)

type IUsecase interface {
	Start(ctx context.Context, testId int, userId int) (*answersheet_struct.StartOutput, error)
	UserAnswer(ctx context.Context, userId int, input answersheet_struct.UserAnswerInput) error
	GetContent(ctx context.Context, testId int, userId int) (*answersheet_struct.GetContentOutput, error)
	CheckUserDoingTest(ctx context.Context, userId int, testId int) (bool, error)
}

type transport struct {
	usecase IUsecase
	answersheetpb.UnimplementedAnswerSheetServiceServer
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) StartDoTest(ctx context.Context, request *answersheetpb.StartDoTestRequest) (*answersheetpb.StartDoTestResponse, error) {
	// !(a || b ) -> !a && !b
	if request.Version != "v1" && request.Version != "v2" {
		panic(errpkg.General.Internal)
	}

	ctx = context.WithValue(ctx, "version", request.Version)
	userId, err := utils.GetAuth(ctx)
	if err != nil {
		panic(err)
	}
	_, err = t.usecase.Start(ctx, int(request.TestId), userId)
	if err != nil {
		panic(err)
	}

	resp := &answersheetpb.StartDoTestResponse{
		Message: "success",
	}

	return resp, nil
}

func (t *transport) UserAnswer(ctx context.Context, request *answersheetpb.UserAnswerRequest) (*answersheetpb.UserAnswerResponse, error) {

	input := answersheet_struct.UserAnswerInput{
		TestId:         int(request.TestId),
		QuestionId:     int(request.QuestionId),
		Answer:         request.Answer,
		PreviousAnswer: request.PreviousAnswer,
	}
	userId, err := utils.GetAuth(ctx)
	if err != nil {
		panic(errpkg.General.Forbidden)
	}
	err = t.usecase.UserAnswer(ctx, userId, input)
	if err != nil {
		panic(err)
	}

	resp := &answersheetpb.UserAnswerResponse{
		Message: "success",
	}

	return resp, nil

}

func (t *transport) GetTestContent(ctx context.Context, request *answersheetpb.GetTestContentRequest) (*answersheetpb.GetTestContentResponse, error) {

	userId, err := utils.GetAuth(ctx)
	if err != nil {
		panic(errpkg.General.Forbidden)
	}
	output, err := t.usecase.GetContent(ctx, int(request.TestId), userId)
	if err != nil {
		panic(err)
	}
	content := output.Content

	data := answersheetpb.TestContent{
		Id:         int64(content.Id),
		TestId:     int64(content.TestId),
		TypeableId: int64(content.TypeableId),
		Typeable:   int64(content.Typeable),
		MultipleChoice: &answersheetpb.TestMultipleChoice{
			Id:       int64(content.MultipleChoice.Id),
			FilePath: content.MultipleChoice.FilePath,
			Score:    float32(content.MultipleChoice.Score),
		},
	}

	if output.TimeLeft != nil {
		left := durationpb.New(*output.TimeLeft)
		data.TimeLeft = left
	}

	answers := make([]*answersheetpb.TestMultipleChoiceAnswer, 0)

	for _, item := range content.MultipleChoice.Answers {
		answers = append(answers, &answersheetpb.TestMultipleChoiceAnswer{
			Id:                   int64(item.Id),
			TestMultipleChoiceId: int64(item.TestMultipleChoiceId),
			Score:                float32(item.Score),
			Type:                 int32(item.Type),
		})
	}
	data.MultipleChoice.Answers = answers

	resp := answersheetpb.GetTestContentResponse{
		Message: "success",
		Data:    &data,
	}
	return &resp, nil
}

func (t *transport) CheckUserDoingTest(ctx context.Context, request *answersheetpb.CheckUserDoingTestRequest) (*answersheetpb.CheckUserDoingTestResponse, error) {

	userId, err := utils.GetAuth(ctx)
	if err != nil {
		panic(err)
	}
	check, err := t.usecase.CheckUserDoingTest(ctx, userId, int(request.TestId))
	if err != nil {
		panic(err)
	}
	resp := answersheetpb.CheckUserDoingTestResponse{
		Check:   check,
		Message: "success",
	}
	return &resp, nil
}

func (t *transport) GetCurrentTest(ctx context.Context, request *answersheetpb.GetCurrentTestRequest) (*answersheetpb.GetCurrentTestResponse, error) {
	userId, err := utils.GetAuth(ctx)
	if err != nil {
		panic(err)
	}
	client, err := grpc.Dial("localhost:30000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	conn := answersheetpb.NewAnswerSheetServiceClient(client)
	resp, err := conn.GetCurrentTest(ctx, &answersheetpb.GetCurrentTestRequest{
		UserId: int64(userId),
		TestId: request.TestId,
	})
	return resp, err
}
