package repository

import (
	"context"
	"math"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

type (
	MobilRepository interface {
		CreateMobil(ctx context.Context, tx *gorm.DB, mobil entity.Mobil) (entity.Mobil, error)
		GetAllMobilWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllMobilRepositoryResponse, error)
		GetMobilById(ctx context.Context, tx *gorm.DB, mobilId string) (entity.Mobil, error)
		UpdateMobil(ctx context.Context, tx *gorm.DB, mobil entity.Mobil) (entity.Mobil, error)
		DeleteMobil(ctx context.Context, tx *gorm.DB, mobilId string) error
	}

	mobilRepository struct {
		db *gorm.DB
	}
)

func NewMobilRepository(db *gorm.DB) MobilRepository {
	return &mobilRepository{
		db: db,
	}
}

func (r *mobilRepository) CreateMobil(ctx context.Context, tx *gorm.DB, mobil entity.Mobil) (entity.Mobil, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&mobil).Error; err != nil {
		return entity.Mobil{}, err
	}

	return mobil, nil
}

func (r *mobilRepository) GetAllMobilWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllMobilRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var mobils []entity.Mobil
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Mobil{}).Count(&count).Error; err != nil {
		return dto.GetAllMobilRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Preload("Merek").Find(&mobils).Error; err != nil {
		return dto.GetAllMobilRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllMobilRepositoryResponse{
		Mobils: mobils,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *mobilRepository) GetMobilById(ctx context.Context, tx *gorm.DB, mobilId string) (entity.Mobil, error) {
	if tx == nil {
		tx = r.db
	}

	var mobil entity.Mobil
	if err := tx.WithContext(ctx).Preload("Merek").Where("id = ?", mobilId).Take(&mobil).Error; err != nil {
		return entity.Mobil{}, err
	}

	return mobil, nil
}

func (r *mobilRepository) UpdateMobil(ctx context.Context, tx *gorm.DB, mobil entity.Mobil) (entity.Mobil, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&mobil).Error; err != nil {
		return entity.Mobil{}, err
	}

	return mobil, nil
}

func (r *mobilRepository) DeleteMobil(ctx context.Context, tx *gorm.DB, mobilId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Mobil{}, "id = ?", mobilId).Error; err != nil {
		return err
	}

	return nil
}
