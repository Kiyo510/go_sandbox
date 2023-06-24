package mysql

import (
	"context"
	"database/sql"

	"github.com/Kiyo510/go_sandbox/pkg/domain/model"
	"github.com/Kiyo510/go_sandbox/pkg/domain/repository"
)

type studentRepository struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) repository.IStudentRepository {
	return &studentRepository{
		DB: db,
	}
}

func (sr *studentRepository) SelectAllStudents(ctx context.Context) (model.StudentSlice, error) {
	// concrete DB operation
	return model.Students().All(ctx, sr.DB)
}
