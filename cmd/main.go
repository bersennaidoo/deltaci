package main

import (
	"flag"
	"log"
	"os"

	"github.com/bersennaidoo/deltaci/internal/app"
	"github.com/bersennaidoo/deltaci/internal/service"
)

func main() {
	proj := flag.String("p", "", "Project directory")
	flag.Parse()

	cis := service.NewCiService()

	a := app.NewApp(cis)

	err := a.ExecuteHandler(*proj, os.Stdout)
	if err != nil {
		log.Println(err)
	}
}
