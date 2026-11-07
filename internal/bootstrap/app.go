package bootstrap

import (
	"context"
	"high-performance-server/internal/config"
	"high-performance-server/internal/httpserver"
	"time"

	"golang.org/x/sync/errgroup"
	"errors"
)

type Application struct{

	CfgProvider config.ConfigProvider
	httpserver *httpserver.HttpServer
}



func NewApplication() (*Application,error){

	cfgprovider :=config.NewLocalConfigProvider()
	cfgprovider.Load()


	httpserver,err := httpserver.NewHttpServer(cfgprovider.Get().HttpConfig)

	if err != nil{
		//add log later on
		return &Application{},err
	}

	return &Application{
		CfgProvider: cfgprovider,
		httpserver: httpserver,
	},nil

}



func (a *Application) Run(ctx context.Context) error{

	workers,workerContext := errgroup.WithContext(ctx)


	workers.Go(func()error{

		return a.httpserver.Run(workerContext)
	})


	workers.Go(func() error{
		<-workerContext.Done()

		shutdowncontext , cancel := context.WithTimeout(context.Background(),10*time.Second)
		defer cancel()
		return a.httpserver.Stop(shutdowncontext)
	})


	err := workers.Wait()

	if errors.Is(err, context.Canceled) {
		return nil
	}
	return err

}


func (a *Application) Stop(){

	

}


