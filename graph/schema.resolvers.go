package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-with-go/graph/generated"
	"graphql-with-go/graph/model"
	"math/rand"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	book := &model.Book{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Title:  input.Title,
		Year:   input.Year,
		Author: &model.Author{ID: input.AuthorID},
	}
	return r.serviceBook.CreateBook(book)
}

func (r *mutationResolver) UpdateBook(ctx context.Context, input model.UpdateBook) (*model.Book, error) {
	book := &model.Book{
		ID:     input.ID,
		Title:  input.Title,
		Year:   input.Year,
		Author: &model.Author{ID: input.AuthorID},
	}
	err := r.serviceBook.UpdateBook(book)

	if err != nil {
		return nil, err
	} else {
		return book, nil
	}
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*model.Message, error) {
	err := r.serviceBook.DeleteBook(id)

	if err != nil {
		return nil, err
	} else {
		return &model.Message{
			Message: "Success",
		}, nil
	}
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return r.serviceBook.ListBook()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
