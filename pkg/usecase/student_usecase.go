package usecase

import (
	"context"

	"github.com/Kiyo510/go_sandbox/pkg/domain/service"
	"github.com/Kiyo510/go_sandbox/pkg/usecase/model"
)

type IStudentUsecase interface {
	FindAllStudents(ctx context.Context) (model.StudentSlice, error)
}

type studentUsecase struct {
	svc service.IStudentService
}

func NewUserUsecase(ss service.IStudentService) IStudentUsecase {
	return &studentUsecase{
		svc: ss,
	}
}

func (su *studentUsecase) FindAllStudents(ctx context.Context) (model.StudentSlice, error) {
	msSlice, err := su.svc.FindAllStudents(ctx)
	if err != nil {
		return nil, err
	}

	sSlice := make(model.StudentSlice, 0, len(msSlice))
	for _, ms := range msSlice {
		sSlice = append(sSlice, model.StudentFromDomainModel(ms))
	}

	return sSlice, nil
}