package store

import (
	"book-server/model"
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type PgBookRepo struct {
	db *DB
}

func NewBookRepo(db *DB) *PgBookRepo {
	return &PgUserRepo{db: db}
}

func (repo *PgBookRepo) GetBook(ctx context.Context, id uuid.UUID) (*model.DbBook, error) {
	book := &model.DbBook{}
	err := repo.db.Model(book).Where("id = ?", id).Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return book, nil
}

// func (repo *PgBookRepo) CreateBook(ctx context.Context, book *model.DbBook) (*model.DBUser, error) {
// 	_, err := repo.db.Model(user).
// 		Returning("*").
// 		Insert()

// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
