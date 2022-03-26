package bookCsvLineMapper

import (
	"github.com/HalukErd/Week4Assignment/domain/author"
	"github.com/HalukErd/Week4Assignment/domain/book"
	"github.com/HalukErd/Week4Assignment/models"
	"github.com/google/uuid"
	"strconv"
)

type BookCsvLineMapper struct {
}

func NewBookCsvLineMapper() *BookCsvLineMapper {
	return &BookCsvLineMapper{}
}

// GetBooksAndAuthors creates Books and Authors from CsvLines
func (mapper *BookCsvLineMapper) GetBooksAndAuthors(lines models.BookCsvLines) (book.Books, author.Authors, error) {
	var books book.Books
	var authors author.Authors
	for _, line := range lines {
		author := author.Author{
			ID:   uuid.New(),
			Name: line.Author,
		}
		pages, err := strconv.Atoi(line.Height)
		if err != nil {
			return nil, nil, err
		}
		book := book.Book{
			ID:        uuid.New(),
			Name:      line.Title,
			Genre:     line.Genre,
			Pages:     pages,
			Publisher: line.Publisher,
			AuthorID:  author.ID,
		}
		authors = append(authors, author)
		books = append(books, book)
	}
	return books, authors, nil
}
