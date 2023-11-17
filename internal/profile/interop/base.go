package interop

import (
	"context"
	"core-crew.runwayclub.dev/core/domain/profile"
)

type ProfileInterop struct {
	profileUseCase profile.ProfileUseCase
}

func (p *ProfileInterop) Create(ctx context.Context, token string, profile *profile.Profile) error {
	return p.profileUseCase.Create(ctx, profile)
}

func (p *ProfileInterop) GetById(ctx context.Context, token string, id string) (*profile.Profile, error) {
	return p.profileUseCase.GetById(ctx, id)
}

func (p *ProfileInterop) Update(ctx context.Context, token string, profile *profile.Profile) error {
	return p.profileUseCase.Update(ctx, profile)
}

func (p *ProfileInterop) Delete(ctx context.Context, token string, id string) error {
	return p.profileUseCase.Delete(ctx, id)
}

func NewProfileInterop(profileUseCase profile.ProfileUseCase) *ProfileInterop {
	return &ProfileInterop{
		profileUseCase: profileUseCase,
	}
}
