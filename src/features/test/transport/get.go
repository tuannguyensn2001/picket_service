package test_transport

import (
	"context"
	testpb "picket/src/pb/test"
	"picket/src/utils"
)

func (t *transport) Get(ctx context.Context, request *testpb.GetTestsRequest) (*testpb.GetTestResponse, error) {

	userId, err := utils.GetAuth(ctx)
	if err != nil {
		panic(err)
	}

	list, err := t.usecase.GetTestsByUserId(ctx, userId)
	if err != nil {
		panic(err)
	}

	result := make([]*testpb.Test, len(list))

	for index, item := range list {
		result[index] = &testpb.Test{
			Id:                 int64(item.Id),
			Name:               item.Name,
			TimeToDo:           int32(item.TimeToDo),
			TimeStart:          utils.ParseTimeToGrpc(item.TimeStart),
			TimeEnd:            utils.ParseTimeToGrpc(item.TimeEnd),
			DoOnce:             item.DoOnce,
			Password:           item.Password,
			PreventCheat:       uint32(item.PreventCheat),
			IsAuthenticateUser: item.IsAuthenticateUser,
			ShowMark:           uint32(item.ShowMark),
			ShowAnswer:         uint32(item.ShowAnswer),
			CreatedBy:          int64(item.CreatedBy),
			CreatedAt:          utils.ParseTimeToGrpc(item.CreatedAt),
			UpdatedAt:          utils.ParseTimeToGrpc(item.UpdatedAt),
		}
	}
	resp := testpb.GetTestResponse{
		Message: "success",
		Data:    result,
	}

	return &resp, nil
}
