package test_transport

import (
	"context"
	testpb "picket/src/pb/test"
	"picket/src/utils"
)

func (t *transport) GetPreview(ctx context.Context, request *testpb.GetTestPreviewRequest) (*testpb.GetTestPreviewResponse, error) {

	test, err := t.usecase.GetPreview(ctx, int(request.Id))
	if err != nil {
		panic(err)
	}

	resp := &testpb.GetTestPreviewResponse{
		Message: "success",
		Data: &testpb.TestPreview{
			Id:        int64(test.Id),
			Name:      test.Name,
			TimeToDo:  int32(test.TimeToDo),
			TimeStart: utils.ParseTimeToGrpc(test.TimeStart),
			TimeEnd:   utils.ParseTimeToGrpc(test.TimeEnd),
			DoOnce:    test.DoOnce,
		},
	}

	return resp, nil
}
