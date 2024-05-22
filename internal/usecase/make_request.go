package usecase

import "renatonasc/stresstest/internal/entity"

type MakeRequestUseCase struct {
	Request entity.Request
}

func NewMakeRequest(request entity.Request) *MakeRequestUseCase {

	return &MakeRequestUseCase{
		Request: request,
	}
}

func (u *MakeRequestUseCase) Execute() entity.RequestResult {
	return entity.RequestResult{}
}
