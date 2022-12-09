package answersheet_transport

import (
	"context"
	answersheet_struct "picket/src/features/answersheet/struct"
	answersheetpb "picket/src/pb/answer_sheet"
	"picket/src/utils"
)

type IUsecase interface {
	Start(ctx context.Context, testId int, userId int) (*answersheet_struct.StartOutput, error)
}

type transport struct {
	usecase IUsecase
	answersheetpb.UnimplementedAnswerSheetServiceServer
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) StartDoTest(ctx context.Context, request *answersheetpb.StartDoTestRequest) (*answersheetpb.StartDoTestResponse, error) {

	userId,err := utils.GetAuth(ctx)
	if err != nil {
		panic(err)
	}
	_,err = t.usecase.Start(ctx, int(request.TestId),userId)
	if err != nil {
		panic(err)
	}

	resp := &answersheetpb.StartDoTestResponse{
		Message: "success",
	}

	return resp,nil
}
