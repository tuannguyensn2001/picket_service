package answersheet_usecase

import (
	"context"
	"picket/src/entities"
)

func (u *usecase) CheckUserDoingTest(ctx context.Context, userId int, testId int) (bool,error)  {
	list,err := u.repository.GetLatestEvent(ctx,userId, testId,2)
	if err != nil {
		return false,err
	}
	if len(list) == 0 {
		return false,nil
	}
	if len(list) == 1  {
		if list[0].Event == entities.START || list[0].Event == entities.DOING {
			return true,nil
		}
		return false,nil
	}

	first,second := list[0],list[1]

	if first.Event == entities.END && second.Event == entities.START {
		return false,nil
	}

	return true,nil
}

