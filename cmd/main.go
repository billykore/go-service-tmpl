package main

import (
	"github.com/billykore/go-service-tmpl/infra/http/server"
	"github.com/billykore/go-service-tmpl/pkg/config"
	_ "github.com/joho/godotenv/autoload"
)

type app struct {
	ss *server.Server
}

func newApp(ss *server.Server) *app {
	return &app{
		ss: ss,
	}
}

func main() {
	c := config.Get()
	a := initApp(c)
	a.ss.Serve()
}
