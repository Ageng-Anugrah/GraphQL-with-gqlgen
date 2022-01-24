package service

import (
	"errors"
	"graphql-with-go/graph/model"
)

type BookServices struct {
	books []*model.Book
}

func (bs *BookServices) CreateBook(book *model.Book) (*model.Book, error) {
	bs.books = append(bs.books, book)
	return book, nil
}

func (bs *BookServices) DeleteBook(id string) error {
	for i := range bs.books {
		if *&bs.books[i].ID == id {
			bs.books = append(bs.books[:i], bs.books[i+1:]...)

			return nil
		}
	}

	return errors.New("cannot find book with given id")
}

func (bs *BookServices) UpdateBook(book *model.Book) error {
	for i := range bs.books {
		if *&bs.books[i].ID == *&book.ID {
			*&bs.books[i] = book

			return nil
		}
	}

	return errors.New("cannot find book with given id")
}

func (bs *BookServices) ListBook() ([]*model.Book, error) {
	return bs.books, nil
}
