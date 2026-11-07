package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"high-performance-server/internal/bootstrap"
)

func main() {


	appctx := context.Background()

	app,err := bootstrap.NewApplication()

	if err != nil{

	}

	appctx,cancel := signal.NotifyContext(appctx,os.Interrupt,syscall.SIGTERM)
	defer cancel()

	
	if err := app.Run(appctx); err !=nil{
		//log error
	}






}
