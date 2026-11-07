package config


type AppConfig struct{

	HttpConfig HttpConfig

}


type HttpConfig struct{

	Port int32
	ReadTimeout int32
	ReadHeaderTimeout int32
	WriteTimeout int32
	IdleTimeout int32

}