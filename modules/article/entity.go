package article

import "time"

type Article struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `validate:"required"`
	Content   string `validate:"required"`
	Image     string `validate:"required, file"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
