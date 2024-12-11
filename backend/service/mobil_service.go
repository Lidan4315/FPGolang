package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/Caknoooo/go-gin-clean-starter/repository"
	"gorm.io/gorm"
)

type MobilService interface {
	CreateMobil(ctx context.Context, mobil entity.Mobil) (entity.Mobil, error)
	GetAllMobil(ctx context.Context, req dto.PaginationRequest) (dto.GetAllMobilRepositoryResponse, error)
	GetMobilById(ctx context.Context, mobilId string) (entity.Mobil, error)
	UpdateMobil(ctx context.Context, mobil entity.Mobil) (entity.Mobil, error)
	DeleteMobil(ctx context.Context, mobilId string) error
}

type mobilService struct {
	repo repository.MobilRepository
	db   *gorm.DB
}

func NewMobilService(repo repository.MobilRepository, db *gorm.DB) MobilService {
	return &mobilService{
		repo: repo,
		db:   db,
	}
}

func (s *mobilService) CreateMobil(ctx context.Context, mobil entity.Mobil) (entity.Mobil, error) {
	return s.repo.CreateMobil(ctx, nil, mobil)
}

func (s *mobilService) GetAllMobil(ctx context.Context, req dto.PaginationRequest) (dto.GetAllMobilRepositoryResponse, error) {
	return s.repo.GetAllMobilWithPagination(ctx, nil, req)
}

func (s *mobilService) GetMobilById(ctx context.Context, mobilId string) (entity.Mobil, error) {
	return s.repo.GetMobilById(ctx, nil, mobilId)
}

func (s *mobilService) UpdateMobil(ctx context.Context, mobil entity.Mobil) (entity.Mobil, error) {
	return s.repo.UpdateMobil(ctx, nil, mobil)
}

func (s *mobilService) DeleteMobil(ctx context.Context, mobilId string) error {
	return s.repo.DeleteMobil(ctx, nil, mobilId)
}
