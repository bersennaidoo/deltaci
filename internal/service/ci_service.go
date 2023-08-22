package service

import (
	"os/exec"

	"github.com/bersennaidoo/deltaci/internal/domain"
)

type CiService struct {
	//Stepp domain.Step
}

func NewCiService() *CiService {
	return &CiService{}
}

func (cis *CiService) Execute(step domain.Step) (string, error) {

	cmd := exec.Command(step.Exe, step.Args...)

	cmd.Dir = step.Proj

	if err := cmd.Run(); err != nil {
		return "", &domain.StepErr{
			Step:  step.Name,
			Msg:   "failed to execute",
			Cause: err,
		}
	}
	return step.Message, nil
}
