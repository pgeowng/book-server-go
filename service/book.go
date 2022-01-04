package service

import "context"

type BookService struct {
	ctx   context.Context
	store *store.Store
}


func NewBookService(ctx context.Context, store *store.Store) *BookService {
	return &BookService{
		ctx: ctx,
		store: store
	}
}

func (s *BookService) GetUser(ctx context.Context, id uuid.UUID) (*model.Book, error) {
	dbBook, err := s.store.Book.GetUser(ctx, id)

	if err != nil {
		return nil, errors.Wrap(err, "BookService:")
	}

	if dbBook == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("Book '%s' not found", id.String()))
	}

	return dbBook.ToJSON(), nil
}