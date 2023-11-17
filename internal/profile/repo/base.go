package repo

import (
	"context"
	"core-crew.runwayclub.dev/core/domain/profile"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func (p *ProfileRepository) Create(ctx context.Context, profile *profile.Profile) error {
	return p.db.WithContext(ctx).Create(profile).Error
}

func (p *ProfileRepository) GetById(ctx context.Context, id string) (*profile.Profile, error) {
	pf := &profile.Profile{}
	tx := p.db.WithContext(ctx).Where("id = ?", id).First(pf)
	return pf, tx.Error
}

func (p *ProfileRepository) Update(ctx context.Context, profile *profile.Profile) error {
	return p.db.WithContext(ctx).Save(profile).Error
}

func (p *ProfileRepository) Delete(ctx context.Context, id string) error {
	return p.db.WithContext(ctx).Where("id = ?", id).Delete(&profile.Profile{}).Error
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	err := db.AutoMigrate(&profile.Profile{})
	if err != nil {
		panic(err)
	}
	return &ProfileRepository{
		db: db,
	}
}
