package author

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthorRepo struct {
	db *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) *AuthorRepo {
	return &AuthorRepo{db: db}
}

func (a *AuthorRepo) Migrations() {
	a.db.AutoMigrate(&Author{})
}

// InsertSampleData inserts data from csv file if there is no such Author
func (a *AuthorRepo) InsertSampleData(Authors Authors) {
	for _, author := range Authors {
		a.db.Where(Author{Name: author.Name}).FirstOrCreate(&author)
	}
}

func (a *AuthorRepo) GetAllAuthorsWithoutBooks() Authors {
	var authors Authors
	a.db.Find(&authors)
	return authors
}

func (a *AuthorRepo) FindByName(name string) Authors {
	var authors Authors
	a.db.Where("name LIKE ? ", "%"+name+"%").Find(&authors)
	return authors
}

func (a *AuthorRepo) GetByID(id uuid.UUID) (*Author, error) {
	var author Author
	result := a.db.First(&author, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &author, nil
}

func (a *AuthorRepo) GetAllAuthorsWithBookInformation() (Authors, error) {
	var authors Authors
	result := a.db.Preload("Books").Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

func (a *AuthorRepo) GetAuthorWithBooks(bookId uuid.UUID) (*Author, error) {
	var author Author
	err := a.db.Joins("JOIN books ON books.author_id=authors.id").
		Where("books.id=?", bookId).
		First(&author).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}
