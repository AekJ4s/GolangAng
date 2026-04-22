package main

import (
	"golangbackend/httpserv"
	"golangbackend/infrastructure"
)

func main() {
	infrastructure.InitConfig()
	infrastructure.InitDB(infrastructure.AppCfg.DB)
	defer infrastructure.DB.Close()

	httpserv.Start()
}
