package profile

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Id        string `json:"id" gorm:"primaryKey,unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ProfileRepository interface {
	Create(ctx context.Context, profile *Profile) error
	GetById(ctx context.Context, id string) (*Profile, error)
	Update(ctx context.Context, profile *Profile) error
	Delete(ctx context.Context, id string) error
}

type ProfileUseCase interface {
	Create(ctx context.Context, profile *Profile) error
	GetById(ctx context.Context, id string) (*Profile, error)
	Update(ctx context.Context, profile *Profile) error
	Delete(ctx context.Context, id string) error
}

type ProfileInterop interface {
	Create(ctx context.Context, token string, profile *Profile) error
	GetById(ctx context.Context, token string, id string) (*Profile, error)
	Update(ctx context.Context, token string, profile *Profile) error
	Delete(ctx context.Context, token string, id string) error
}

var (
	ErrInvalidId = errors.New("invalid id")
)
