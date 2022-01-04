package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Author    string    `json:"author" validate:"required"`
	Pages     int       `json:"pages"`
	Price     string    `json:"price"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

type DbBook struct {
	tableName struct{}  `pg:"books" gorm:"primaryKey"`
	ID        uuid.UUID `pg:"id,notnull,pk"`
	Title     string    `pg:"title,notnull"`
	Author    string    `pg:"author,notnull"`
	Pages     int       `pg:"pages,notnull"`
	Price     string    `pg:"price"`
	Rating    int       `pg:"rating"`
	CreatedAt time.Time `pg:"created_at,notnull"`
}

func (DbBook) TableName() string {
	return "books"
}

func (b *Book) ToDB() *DbBook {
	return &DbBook{
		ID: b.ID,
		Title: b.Title,
		Author: b.Author,
		Pages: b.Pages,
		Price: b.Price,
		Rating: b.Rating,
		CreatedAt: b.CreatedAt
	}
}

func (b *DbBook) ToJSON() *Book {
	return &Book {
		ID: b.ID,
		Title: b.Title,
		Author: b.Author,
		Pages: b.Pages,
		Price: b.Price,
		Rating: b.Rating,
		CreatedAt: b.CreatedAt
	}
}
