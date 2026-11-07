package config

import (
	"context"
)


type ConfigProvider interface{

	Load()
	Watch(ctx context.Context)
	Get() *AppConfig
}