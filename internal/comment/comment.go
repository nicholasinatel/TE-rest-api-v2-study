package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - defines all of the methods that our service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	// UpdateComment(context.Context, Comment) error
	// DeleteComment(context.Context, string) error
	// CreateComment(context.Context, Comment) (Comment, error)
}

// Service - is the struct on which all our logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	// Will accept any defined type of Store
	return &Service{
		Store: store,
	}
}

// GetComment - method that returns comment throw service implemented by the Store interface
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")

	// Decoupled Business Layer
	// Perfect for testing without touching the database with mocked responses
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

// func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
// 	return ErrNotImplemented
// }

// func (s *Service) DeleteComment(ctx context.Context, id string) error {
// 	return ErrNotImplemented
// }

// func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
// 	return Comment{}, ErrNotImplemented
// }
