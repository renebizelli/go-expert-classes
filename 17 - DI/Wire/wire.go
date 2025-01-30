//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

var setRepoInterface = wire.NewSet(
	NewServiceRepo,
	wire.Bind(new(IGetRepo), new(*ServiceRepository)),
)

func NewServiceWire() *Service {
	wire.Build(setRepoInterface, NewService)

	return &Service{}
}
