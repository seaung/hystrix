package models

import "time"

type DomainTaskModel struct {
	ID          int64     `gorm:"colunm:id;primary_key"`
	TaskID      int64     `gorm:"colunm:task_id"`
	TaskName    string    `gorm:"colunm:task_name;not null"`
	Status      string    `gorm:"colunm:status"`
	CreatedTime time.Time `gorm:"colunm:created_time"`
	UpdatedTime time.Time `gorm:"colunm:updated_time"`
}

func (d *DomainTaskModel) TableName() string {
	return "domain_task"
}
