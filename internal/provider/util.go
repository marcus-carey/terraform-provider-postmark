package provider

import (
	"context"
	"time"
)

func GetRequestContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	return ctx
}
