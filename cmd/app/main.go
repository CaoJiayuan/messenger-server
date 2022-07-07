package main

import (
	"log"
	"os"

	"messenger-server/internal/app"
	"messenger-server/internal/pkg/env"
	"messenger-server/internal/pkg/path"
	"github.com/enorith/framework"
)

func main() {
	// load .env, before app created
	env.LoadDotenv()
	application := framework.NewApp(os.DirFS(path.BasePath("config")), path.BasePath("storage/logs"))
	app.BootstrapApp(application)

	e := application.Run()
	if e != nil {
		log.Fatal(e)
	}
}
