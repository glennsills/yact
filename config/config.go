package config

import (
	"flag"
	"log"
	"os"
)

var App struct {
	Addr string
	Env  string
	Dsn  string
	Log  *log.Logger
}

func InitAppConfig() {
	flag.StringVar(&App.Addr, "port", ":4000", "Http Network Address")
	flag.StringVar(&App.Env, "env", "dev", "Environment (dev | stage | prod)")
	flag.StringVar(&App.Dsn, "db-dsn", os.Getenv("GORM_DSN"), "PostgreSQL DSN")
	flag.Parse()
	App.Log = log.New(os.Stdout, "", log.Ldate|log.Ltime)
}
