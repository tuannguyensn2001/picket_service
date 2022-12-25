package answersheet_usecase

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	answersheetpb "picket/src/pb/answer_sheet"
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
	resp, err := conn.CheckUserDoingTest(ctx, &answersheetpb.CheckUserDoingTestRequest{
		UserId: int64(userId),
		TestId: int64(testId),
	})
	if err != nil {
		return false, err
	}
	return resp.Check, nil
}
