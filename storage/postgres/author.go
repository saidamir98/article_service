package postgres

import (
	"fmt"
	"time"
	"uacademy/blogpost/article_service/protogen/blogpost"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetAuthorByID ...
func (stg Postgres) GetAuthorByID(id string) (*blogpost.Author, error) {
	result := &blogpost.Author{}
	var createdAt, updatedAt *time.Time
	err := stg.db.QueryRow(`SELECT 
		id,
		fullname,
		created_at,
		updated_at
    FROM author WHERE id = $1`, id).Scan(
		&result.Id,
		&result.Fullname,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return result, err
	}

	if createdAt != nil {
		result.CreatedAt = timestamppb.New(*createdAt)
	}

	if updatedAt != nil {
		result.UpdatedAt = timestamppb.New(*updatedAt)
	}

	fmt.Printf("%+v", result)

	return result, nil
}
