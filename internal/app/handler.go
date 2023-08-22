package app

import (
	"fmt"
	"io"

	"github.com/bersennaidoo/deltaci/internal/domain"
)

func (a *App) ExecuteHandler(proj string, out io.Writer) error {

	pipeline := make([]domain.Step, 1)

	pipeline[0] = domain.NewStep("go build", "go", "Go Build: SUCCESS",
		proj, []string{"build", "."})

	for _, s := range pipeline {
		msg, err := a.Cis.Execute(s)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
