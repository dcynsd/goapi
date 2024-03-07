package services

import "goapi/app/models"

type BaseService struct {
	Error  *models.Error
	UserID uint64
}
