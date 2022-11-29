package postgres

import (
	"uacademy/blogpost/article_service/models"
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

	err := stg.db.QueryRow(`SELECT 
		id,
		fullname,
		created_at,
		updated_at,
		deleted_at
    FROM author WHERE id = $1`, id).Scan(
		&result.ID,
		&result.Fullname,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetAuthorList ...
func (stg Postgres) GetAuthorList() (resp []models.Author, err error) {
	// resp = im.Db.InMemoryAuthorData
	return resp, err
}
