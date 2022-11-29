package postgres

import (
	"errors"
	"uacademy/article/models"
)

// AddAuthor ...
func (stg Postgres) AddAuthor(id string, entity models.CreateAuthorModel) error {
	// var author models.Author
	// author.ID = id
	// author.Firstname = entity.Firstname
	// author.Lastname = entity.Lastname
	// author.CreatedAt = time.Now()

	// im.Db.InMemoryAuthorData = append(im.Db.InMemoryAuthorData, author)
	return nil
}

// GetAuthorByID ...
func (stg Postgres) GetAuthorByID(id string) (models.Author, error) {
	var result models.Author
	// for _, v := range im.Db.InMemoryAuthorData {
	// 	if v.ID == id {
	// 		result = v
	// 		return result, nil
	// 	}
	// }
	return result, errors.New("author not found")
}

// GetAuthorList ...
func (stg Postgres) GetAuthorList() (resp []models.Author, err error) {
	// resp = im.Db.InMemoryAuthorData
	return resp, err
}
