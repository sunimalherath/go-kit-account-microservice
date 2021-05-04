package account

import "github.com/go-kit/kit/endpoint"

type Endpoint struct {
	CreateUser endpoint.Endpoint
	GetUser endpoint.Endpoint
}
