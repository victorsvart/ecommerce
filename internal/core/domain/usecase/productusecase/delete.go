package productusecase

import "context"

func (p *productUseCase) Delete(ctx context.Context, id uint64) error {
	if err := p.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
