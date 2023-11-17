package usecase

import (
	"context"
	"core-crew.runwayclub.dev/core/domain/profile"
)

type ProfileUseCase struct {
	repo profile.ProfileUseCase
}

func (p *ProfileUseCase) Create(ctx context.Context, profile *profile.Profile) error {
	return p.repo.Create(ctx, profile)
}

func (p *ProfileUseCase) GetById(ctx context.Context, id string) (*profile.Profile, error) {
	if id == "" {
		return nil, profile.ErrInvalidId
	}
	return p.repo.GetById(ctx, id)
}

func (p *ProfileUseCase) Update(ctx context.Context, profile *profile.Profile) error {
	return p.repo.Update(ctx, profile)
}

func (p *ProfileUseCase) Delete(ctx context.Context, id string) error {
	return p.repo.Delete(ctx, id)
}

func NewProfileUseCase(repo profile.ProfileUseCase) *ProfileUseCase {
	return &ProfileUseCase{
		repo: repo,
	}
}
