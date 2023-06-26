package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

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

func (sr *studentRepository) SelectStudentByID(ctx context.Context, studentId int) (*model.Student, error) {
	whereId := fmt.Sprintf("%s = ?", model.StudentColumns.ID)
	return model.Students(qm.Where(whereId, studentId)).One(ctx, sr.DB)
}
