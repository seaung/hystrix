package store

import (
	"context"
	"errors"

	"github.com/seaung/hystrix/internal/pkg/models"
	"gorm.io/gorm"
)

type DomainTaskStore interface {
	CreateDomainTask(ctx context.Context, task *models.DomainTaskModel) error
	GetDomainTask(ctx context.Context, taskName string) (*models.DomainTaskModel, error)
	UpdateDomainTaskStatus(ctx context.Context, taskID int64, status string) error
	ListDomainTasks(ctx context.Context, offset, limit int) (int64, []*models.DomainTaskModel, error)
	DeleteDomainTaks(ctx context.Context, TaskID int64) error
}

type tasks struct {
	db *gorm.DB
}

var _ DomainTaskStore = (*tasks)(nil)

func newTasks(db *gorm.DB) *tasks {
	return &tasks{db}
}

func (t *tasks) CreateDomainTask(ctx context.Context, task *models.DomainTaskModel) error {
	return t.db.Create(&task).Error
}

func (t *tasks) GetDomainTask(ctx context.Context, taskName string) (*models.DomainTaskModel, error) {
	var task models.DomainTaskModel
	if err := t.db.Where("TaskID = ?", taskName).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *tasks) UpdateDomainTaskStatus(ctx context.Context, taskID int64, status string) error {
	err := t.db.Where("TaskID = ?", taskID).Update("status", status).Error
	return err
}

func (t *tasks) ListDomainTasks(ctx context.Context, offset, limit int) (count int64, res []*models.DomainTaskModel, err error) {
	err = t.db.Offset(offset).Limit(20).Order("id desc").Find(&res).Offset(-1).Limit(-1).Count(&count).Error
	return
}

func (t *tasks) DeleteDomainTaks(ctx context.Context, TaskID int64) error {
	err := t.db.Where("TaskID = ?").Delete(&models.DomainTaskModel{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
