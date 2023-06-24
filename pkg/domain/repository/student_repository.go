package repository

import (
	"context"

	"github.com/Kiyo510/go_sandbox/pkg/domain/model"
)

// IHogeHoge represents interface of HogeHoge
type IStudentRepository interface {
	SelectAllStudents(ctx context.Context) (model.StudentSlice, error)
}
