package app

import "github.com/bersennaidoo/deltaci/internal/service"

type App struct {
	Cis *service.CiService
}

func NewApp(cis *service.CiService) *App {
	return &App{
		Cis: cis,
	}
}
