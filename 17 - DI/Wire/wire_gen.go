// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewServiceWire() *Service {
	serviceRepository := NewServiceRepo()
	service := NewService(serviceRepository)
	return service
}

// wire.go:

var setRepoInterface = wire.NewSet(
	NewServiceRepo, wire.Bind(new(IGetRepo), new(*ServiceRepository)),
)
