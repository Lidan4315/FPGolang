package dto

import (
	"errors"
	"mime/multipart"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
)

const (
	// Failed
	MESSAGE_FAILED_CREATE_MOBIL      = "failed create mobil"
	MESSAGE_FAILED_GET_LIST_MOBIL    = "failed get list mobil"
	MESSAGE_FAILED_GET_MOBIL         = "failed get mobil"
	MESSAGE_FAILED_UPDATE_MOBIL      = "failed update mobil"
	MESSAGE_FAILED_DELETE_MOBIL      = "failed delete mobil"
	MESSAGE_FAILED_PROSES_REQUEST_MOBIL    = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS_MOBIL     = "denied access"

	// Success
	MESSAGE_SUCCESS_CREATE_MOBIL     = "success create mobil"
	MESSAGE_SUCCESS_GET_LIST_MOBIL   = "success get list mobil"
	MESSAGE_SUCCESS_GET_MOBIL        = "success get mobil"
	MESSAGE_SUCCESS_UPDATE_MOBIL     = "success update mobil"
	MESSAGE_SUCCESS_DELETE_MOBIL     = "success delete mobil"
)

var (
	ErrCreateMobil        = errors.New("failed to create mobil")
	ErrGetAllMobil        = errors.New("failed to get all mobil")
	ErrGetMobilById       = errors.New("failed to get mobil by id")
	ErrUpdateMobil        = errors.New("failed to update mobil")
	ErrDeleteMobil        = errors.New("failed to delete mobil")
	ErrMobilNotFound      = errors.New("mobil not found")
)

type (
	// Create request for a new Mobil
	MobilCreateRequest struct {
		MerekID   string                `json:"merek_id" form:"merek_id"`
		Type      string                `json:"type" form:"type"`
		NoPlat    string                `json:"no_plat" form:"no_plat"`
		Warna     string                `json:"warna" form:"warna"`
		Condition string                `json:"initial_condition" form:"initial_condition"`
		Harga     float64               `json:"harga" form:"harga"`
		Deskripsi string                `json:"deskripsi" form:"deskripsi"`
		Image     *multipart.FileHeader `json:"image" form:"image"`
	}

	// Response for a single mobil
	MobilResponse struct {
		ID               string  `json:"id"`
		Type             string  `json:"type"`
		NoPlat           string  `json:"no_plat"`
		Warna            string  `json:"warna"`
		InitialCondition string  `json:"initial_condition"`
		Harga            float64 `json:"harga"`
		Deskripsi        string  `json:"deskripsi"`
		ImageUrl         string  `json:"image_url"`
		Merek            string  `json:"merek"`  // Nama merek saja, bukan objek lengkap
	}

	// Response for mobil list with pagination
	MobilPaginationResponse struct {
		Data []MobilResponse `json:"data"`
		PaginationResponse
	}

	// Repository response for all mobil
	GetAllMobilRepositoryResponse struct {
		Mobils []entity.Mobil
		PaginationResponse
	}

	// Request for updating an existing mobil
	MobilUpdateRequest struct {
		ID              string  `json:"id" form:"id"`
		MerekID         string  `json:"merek_id" form:"merek_id"`
		Type            string  `json:"type" form:"type"`
		NoPlat          string  `json:"no_plat" form:"no_plat"`
		Warna           string  `json:"warna" form:"warna"`
		InitialCondition string  `json:"initial_condition" form:"initial_condition"`
		Harga           float64 `json:"harga" form:"harga"`
		Deskripsi       string  `json:"deskripsi" form:"deskripsi"`
		Image           *multipart.FileHeader `json:"image" form:"image"`
	}

	// Response after update
	MobilUpdateResponse struct {
		ID              string  `json:"id"`
		MerekID         string  `json:"merek_id"`
		Merek           string  `json:"merek"`       // Merek name
		Type            string  `json:"type"`
		NoPlat          string  `json:"no_plat"`
		Warna           string  `json:"warna"`
		InitialCondition string  `json:"initial_condition"`
		Harga           float64 `json:"harga"`
		Deskripsi       string  `json:"deskripsi"`
		ImageUrl        string  `json:"image_url"`
	}

	// Request for deleting a mobil
	MobilDeleteRequest struct {
		ID string `json:"id" form:"id" binding:"required"`
	}
)
