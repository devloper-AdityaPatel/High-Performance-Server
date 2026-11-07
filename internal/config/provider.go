package config

import "context"


type LocalConfigProvider struct{
	config *AppConfig
}



func NewLocalConfigProvider() *LocalConfigProvider{

	return &LocalConfigProvider{
		config : nil,
	}
}


func (lcp *LocalConfigProvider) Load(){

	lcp.config = &AppConfig{

		HttpConfig: HttpConfig{
			Port:              8080,
			ReadTimeout:       5,
			ReadHeaderTimeout: 5,
			WriteTimeout:      5,
			IdleTimeout:       5,
		},

	}
}

func (lcp *LocalConfigProvider) Watch(ctx context.Context) {

	lcp.config = &AppConfig{

		HttpConfig: HttpConfig{
			Port:              8080,
			ReadTimeout:       5,
			ReadHeaderTimeout: 5,
			WriteTimeout:      5,
			IdleTimeout:       5,
		},

	}
}


func (lcp *LocalConfigProvider) Get() *AppConfig{

	return lcp.config
}