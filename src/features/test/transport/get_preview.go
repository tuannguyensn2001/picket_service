package test_transport

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	testpb "picket/src/pb/test"
)

func (t *transport) GetPreview(ctx context.Context, request *testpb.GetTestPreviewRequest) (*testpb.GetTestPreviewResponse, error) {

	test, err := t.usecase.GetPreview(ctx, request.Code)
	if err != nil {
		panic(err)
	}

	resp := &testpb.GetTestPreviewResponse{
		Message: "success",
		Data: &testpb.TestPreview{
			Id:        int64(test.Id),
			Name:      test.Name,
			TimeToDo:  int32(test.TimeToDo),
			TimeStart: timestamppb.New(*test.TimeStart),
			TimeEnd:   timestamppb.New(*test.TimeEnd),
			DoOnce:    test.DoOnce,
		},
	}

	return resp, nil
}
