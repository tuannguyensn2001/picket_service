package test_usecase

import (
	"context"
	test_struct "myclass_service/src/features/test/struct"
)

func (u *usecase) UpdateContent(ctx context.Context, input test_struct.UpdateTestContentInput) error {

	content, err := u.repository.FindContentByTestId(ctx, input.TestId)
	if err != nil {
		return err
	}
	if content.Typeable == input.Typeable {
		var handler func(ctx context.Context, input test_struct.UpdateTestContentInput) error
		switch content.Typeable {
		case 1:
			handler = u.UpdateMultipleChoiceContent
		}
		err = handler(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *usecase) UpdateMultipleChoiceContent(ctx context.Context, input test_struct.UpdateTestContentInput) error {

	return nil
}
