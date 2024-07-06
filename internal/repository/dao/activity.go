package dao

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ActivityDAO interface {
	GetRecentActivity(ctx context.Context) (RecentActivity, error)
	SetRecentActivity(ctx context.Context, mr RecentActivity) error
}

type activityDAO struct {
	db *gorm.DB
	l  *zap.Logger
}

type RecentActivity struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	UserID      int64  `gorm:"column:user_id;not null" json:"user_id"`
	Description string `gorm:"type:varchar(255);not null"`
	Time        string `gorm:"type:varchar(255);not null"`
}

func NewActivityDAO(db *gorm.DB, l *zap.Logger) ActivityDAO {
	return &activityDAO{
		db: db,
		l:  l,
	}
}

func (a *activityDAO) GetRecentActivity(ctx context.Context) (RecentActivity, error) {
	var mr RecentActivity
	if err := a.db.WithContext(ctx).Model(RecentActivity{}).Find(&mr).Error; err != nil {
		a.l.Error("get recent activity failed", zap.Error(err))
		return RecentActivity{}, err
	}
	return mr, nil
}

func (a *activityDAO) SetRecentActivity(ctx context.Context, mr RecentActivity) error {
	if err := a.db.WithContext(ctx).Create(&mr).Error; err != nil {
		a.l.Error("set recent activity failed", zap.Error(err))
		return err
	}
	return nil
}
