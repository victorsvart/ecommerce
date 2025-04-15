package userusecase

import "context"

func (u *userUseCaseImpl) Delete(ctx context.Context, id uint64) error {
	err := u.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
