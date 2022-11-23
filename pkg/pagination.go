package pkg

import (
	"gorm.io/gorm"
)

const (
	DefaultPage = 1
	DefaultSize = 50
)

type Pagination[T any] struct {
	Items []*T `json:"items"`
	Page  int  `json:"page"`
	Size  int  `json:"size"`
	Total int  `json:"total"`
}

func (p *Pagination[T]) GetPage() int {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}

	return p.Page
}

func (p *Pagination[T]) GetSize() int {
	if p.Size <= 0 {
		p.Size = DefaultSize
	}

	return p.Size
}

func (p *Pagination[T]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

func (p *Pagination[T]) Query(db *gorm.DB) *gorm.DB {
	return db.Offset(p.GetOffset()).Limit(p.GetSize())
}

func (p *Pagination[T]) Paginate(items []*T, total int) *Pagination[T] {
	p.Items = items
	p.Total = total

	return p
}
