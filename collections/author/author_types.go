package author

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID        uuid.UUID  `json:"id"`
	FullName  string     `json:"fullName"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type CreateAuthorRequest struct {
	FullName string `json:"fullName"`
}

type UpdateAuthorRequest struct {
	FullName string `json:"fullName"`
}

type DeleteAuthorResponse struct {
	Message string `json:"message"`
}

func NewAuthor(fullName string) (*Author, error) {
	uuid := uuid.New()
	return &Author{
		ID:        uuid,
		FullName:  fullName,
		CreatedAt: time.Now().UTC(),
	}, nil

}

func UpdateAuthor(fullName string) (*Author, error) {
	time := time.Now().UTC()
	return &Author{
		FullName:  fullName,
		UpdatedAt: &time,
	}, nil
}

func DeleteAuthor() (*Author, error) {
	time := time.Now().UTC()
	return &Author{
		DeletedAt: &time,
	}, nil
}
