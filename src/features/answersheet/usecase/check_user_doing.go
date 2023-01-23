package answersheet_usecase

import (
	"context"
	"errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	answersheetpb "picket/src/pb/answer_sheet"
	"time"
)

func (u *usecase) CheckUserDoingTest(ctx context.Context, userId int, testId int) (bool, error) {
	//list,err := u.repository.GetLatestEvent(ctx,userId, testId,2)
	//if err != nil {
	//	return false,err
	//}
	//if len(list) == 0 {
	//	return false,nil
	//}
	//if len(list) == 1  {
	//	if list[0].Event == entities.START || list[0].Event == entities.DOING {
	//		return true,nil
	//	}
	//	return false,nil
	//}
	//
	//first,second := list[0],list[1]
	//
	//if first.Event == entities.END && second.Event == entities.START {
	//	return false,nil
	//}
	//
	//return true,nil

	client, err := grpc.Dial("localhost:30000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	if err != nil {
		return false, err
	}
	defer client.Close()
	ctx, span := tracer.Start(ctx, "call grpc", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	conn := answersheetpb.NewAnswerSheetServiceClient(client)
	respGetLatestStartTime, err := conn.GetLatestStartTime(ctx, &answersheetpb.GetLatestStartTimeRequest{
		TestId: int64(testId),
		UserId: int64(userId),
	})
	if err != nil {
		return false, nil
	}
	if respGetLatestStartTime.Data != nil {
		t := respGetLatestStartTime.Data.AsTime()
		test, err := u.testUsecase.GetById(ctx, testId)
		if err != nil {
			return false, err
		}
		if test.TimeEnd != nil && test.TimeEnd.Before(time.Now()) {
			return false, nil
		}
		if t.Add(time.Duration(test.TimeToDo) * time.Minute).Before(time.Now()) {
			return false, nil
		}
	}

	resp, err := conn.CheckUserDoingTest(ctx, &answersheetpb.CheckUserDoingTestRequest{
		UserId: int64(userId),
		TestId: int64(testId),
	})
	if err != nil {
		return false, err
	}
	return resp.Check, nil
}

func (u *usecase) GetLatestStartTime(ctx context.Context, testId int, userId int) (*time.Time, error) {
	client, err := grpc.Dial("localhost:30000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer client.Close()
	conn := answersheetpb.NewAnswerSheetServiceClient(client)
	respGetLatestStartTime, err := conn.GetLatestStartTime(ctx, &answersheetpb.GetLatestStartTimeRequest{
		TestId: int64(testId),
		UserId: int64(userId),
	})
	if err != nil {
		return nil, err
	}
	if respGetLatestStartTime.Data != nil {
		result := respGetLatestStartTime.Data.AsTime()
		return &result, nil
	}
	return nil, errors.New("user hasn't start test")
}
